package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"maps"
	"math"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

func averageColor(img image.Image) [3]float64 {
	bounds := img.Bounds()
	r, g, b := 0.0, 0.0, 0.0

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r1, g1, b1, _ := img.At(x, y).RGBA()
			r, g, b = r+float64(r1), g+float64(g1), b+float64(b1)
		}
	}
	totalPixel := float64(bounds.Max.X * bounds.Max.Y)
	return [3]float64{r / totalPixel, g / totalPixel, b / totalPixel}
}

func resize(in image.Image, newWidth int) image.NRGBA {
	bounds := in.Bounds()
	ratio := bounds.Dx() / newWidth
	out := image.NewNRGBA(
		image.Rect(
			bounds.Min.X/ratio,
			bounds.Min.X/ratio,
			bounds.Max.X/ratio,
			bounds.Max.Y/ratio,
		),
	)

	for y, j := bounds.Min.Y, bounds.Min.Y; y < bounds.Max.Y; y, j = y+ratio, j+1 {
		for x, i := bounds.Min.X, bounds.Min.X; x < bounds.Max.X; x, i = x+ratio, x+1 {
			r, g, b, a := in.At(x, y).RGBA()
			out.SetNRGBA(i, j, color.NRGBA{
				uint8(r >> 8),
				uint8(g >> 8),
				uint8(b >> 8),
				uint8(a >> 8),
			})
		}
	}

	return *out
}

func tilesDB() map[string][3]float64 {
	fmt.Println("start populating tiles DB")
	db := make(map[string][3]float64)
	files, err := os.ReadDir("tiles")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		name := "tiles/" + f.Name()
		fmt.Println("Opening tile:", name) // Tambahkan log ini

		file, err := os.Open(name)
		if err != nil {
			fmt.Println("Error opening file:", name, err)
			continue
		}
		defer file.Close()

		img, _, err := image.Decode(file)
		if err != nil {
			fmt.Println("Error decoding image:", name, err)
			continue
		}

		db[name] = averageColor(img)
	}

	fmt.Println("finished populating tiles DB")
	return db
}

func nearest(target [3]float64, db *map[string][3]float64) string {
	var filename string
	smallest := 1_000_000.0
	for k, v := range *db {
		dist := distance(target, v)
		if dist < smallest {
			filename, smallest = k, dist
		}
	}
	if filename == "" {
		fmt.Println("WARNING: nearest() returned empty string")
	}
	return filename
}

// func nearest(target [3]float64, db *map[string][3]float64) string {
// 	var filename string
//
// 	smallest := 1_000_000.0
// 	for k, v := range *db {
// 		dist := distance(target, v)
// 		if dist < smallest {
// 			filename, smallest = k, dist
// 		}
//
// 	}
// 	delete(*db, filename)
// 	return filename
// }

func distance(p1 [3]float64, p2 [3]float64) float64 {
	return math.Sqrt(sq(p2[0]-p1[0]) + sq(p2[1]-p1[1]) + sq(p2[2]-p1[2]))
}

func sq(n float64) float64 {
	return n * n
}

var TILESDB map[string][3]float64

func cloneTilesDB() map[string][3]float64 {
	db := make(map[string][3]float64)
	maps.Copy(db, TILESDB)

	return db
}

func main() {
	mux := http.NewServeMux()
	// files := http.FileServer(http.Dir("public"))
	// mux.Handle("/static", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", upload)
	mux.HandleFunc("/mosaic", mosaic)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	TILESDB = tilesDB()
	fmt.Println("Mosaic server started")
	server.ListenAndServe()
}

func upload(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("upload.html")
	t.Execute(w, nil)
}

func mosaic(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now()

	r.ParseMultipartForm(10485760)
	file, _, _ := r.FormFile("image")
	defer file.Close()

	tileSize, _ := strconv.Atoi(r.FormValue("tile_size"))

	original, _, _ := image.Decode(file)
	bounds := original.Bounds()

	newimage := image.NewRGBA(
		image.Rect(
			bounds.Min.X,
			bounds.Min.X,
			bounds.Max.X,
			bounds.Max.Y,
		),
	)

	db := cloneTilesDB()

	sp := image.Point{0, 0}

	for y := bounds.Min.Y; y < bounds.Max.Y; y = y + tileSize {
		for x := bounds.Min.X; x < bounds.Max.X; x = x + tileSize {

			r, g, b, _ := original.At(x, y).RGBA()
			color := [3]float64{
				float64(r),
				float64(g),
				float64(b),
			}

			nearest := nearest(color, &db)
			file, err := os.Open(nearest)
			if err == nil {
				img, _, err := image.Decode(file)
				if err == nil {
					t := resize(img, tileSize)
					tile := t.SubImage(t.Bounds())
					tileBounds := image.Rect(x, y, x+tileSize, y+tileSize)
					draw.Draw(newimage, tileBounds, tile, sp, draw.Src)
				} else {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}

			file.Close()
		}
	}

	buf1 := new(bytes.Buffer)
	jpeg.Encode(buf1, original, nil)
	originalSTR := base64.StdEncoding.EncodeToString(buf1.Bytes())

	buf2 := new(bytes.Buffer)
	jpeg.Encode(buf2, newimage, nil)
	mosaic := base64.StdEncoding.EncodeToString(buf2.Bytes())
	t1 := time.Now()

	images := map[string]string{
		"original": originalSTR,
		"mosaic":   mosaic,
		"duration": fmt.Sprintf("%v", t1.Sub(t0)),
	}

	t, _ := template.ParseFiles("results.html")
	t.Execute(w, images)
}
