package concurrent

import (
	"log"
	"path/filepath"
)

type PathData struct {
	path string
	err  error
}

func orDone(done <-chan struct{}, c <-chan PathData) <-chan PathData {
	valStream := make(chan PathData)

	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}

				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()

	return valStream
}

func pathGenerator(done <-chan struct{}, path ...string) <-chan PathData {
	stringStream := make(chan PathData)

	go func() {
		defer close(stringStream)

		for _, p := range path {
			path, pathErr := filepath.Abs(p)

			select {
			case <-done:
				return
			case stringStream <- PathData{path: path, err: pathErr}:
			}
		}
	}()

	return stringStream
}

func groupingPath(done <-chan struct{}, pathStream <-chan PathData) <-chan []PathData {
	out := make(chan []PathData, 1)

	go func() {
		defer close(out)

		var arr []PathData

		for p := range orDone(done, pathStream) {
			arr = append(arr, p)
		}

		select {
		case <-done:
			return
		case out <- arr:
		}
	}()

	return out
}

func extractChanPath(data <-chan []PathData) ([]string, error) {
	var files []string

	path, ok := <-data
	if !ok {
		return nil, nil
	}

	for _, v := range path {

		if v.err != nil {
			return nil, v.err
		}

		files = append(files, v.path)
	}
	return files, nil
}

func PathFile(path ...string) ([]string, error) {
	done := make(chan struct{})

	defer close(done)

	data, err := extractChanPath(groupingPath(done, pathGenerator(done, path...)))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return data, nil
}
