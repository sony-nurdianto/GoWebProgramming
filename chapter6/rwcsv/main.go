package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	csvFile, err := os.Create("post.csv")
	if err != nil {
		panic(err)
	}

	defer csvFile.Close()

	allPost := []Post{
		{Id: 1, Content: "Hello world!", Author: "Sau Sheong"},
		{Id: 2, Content: "Halo Welt", Author: "Sony Nurdianto"},
		{Id: 3, Content: "Halo Mundo", Author: "Pedro"},
		{Id: 4, Content: "Greeting Earthlings!", Author: "Sau Sheong"},
	}

	writer := csv.NewWriter(csvFile)
	for _, post := range allPost {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		if err := writer.Write(line); err != nil {
			panic(err)
		}
	}

	writer.Flush()

	file, err := os.Open("post.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post

	for _, item := range record {
		id, _ := strconv.Atoi(item[0])
		post := Post{Id: id, Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}

	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
