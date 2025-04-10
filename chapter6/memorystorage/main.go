package main

import "fmt"

type Post struct {
	Id      int
	Content string
	Author  string
}

var (
	PostById     map[int]*Post
	PostByAuthor map[string][]*Post
)

func store(post Post) {
	PostById[post.Id] = &post
	PostByAuthor[post.Author] = append(PostByAuthor[post.Author], &post)
}

func main() {
	PostById = make(map[int]*Post)
	PostByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "Hello, World!", Author: "Sau Sheong"}
	post2 := Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"}
	post3 := Post{Id: 3, Content: "Hola Mundo", Author: "Pedro"}
	post4 := Post{Id: 4, Content: "Halo Dunia!", Author: "Sony Nurdianto"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, sau := range PostByAuthor["Sau Sheong"] {
		fmt.Println(sau)
	}

	for _, sony := range PostByAuthor["Sony Nurdianto"] {
		fmt.Println(sony)
	}
}
