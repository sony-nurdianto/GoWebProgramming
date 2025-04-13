package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func store[T any](data T, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)

	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filename, buffer.Bytes(), 0666)
	if err != nil {
		panic(err)
	}
}

func load[T any](data *T, filename string) {
	raw, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	post := Post{Id: 1, Content: "Hallo World!", Author: "Sony Nurdianto"}
	post1 := Post{Id: 2, Content: "Halo Welt!", Author: "Sony Nurdianto"}

	store(post, "post")
	store(post1, "post")

	var postRead Post
	var postRead1 Post
	load(&postRead, "post")
	load(&postRead1, "post")

	fmt.Println(postRead)
	fmt.Println(postRead1)
}
