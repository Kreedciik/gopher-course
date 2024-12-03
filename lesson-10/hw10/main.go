package main

import (
	"fmt"
	fs "hw10/files"
)

func main() {
	comments := fs.JsonToStruct("comments.json")
	fmt.Println(comments)
}
