package package2

import (
	"encoding/json"
	p1 "hw11/package1"
	"os"
	"sync"
)

const (
	FILE_NAME = "books.json"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
func WriteToJSON(books chan p1.Book, wg *sync.WaitGroup) {
	book := <-books
	file, e := os.OpenFile(FILE_NAME, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	handleError(e)
	defer file.Close()
	data, err := os.ReadFile(FILE_NAME)
	handleError(err)
	if len(data) > 0 {
		data = append(data[:len(data)-1], ',', '\n')
		encoded, e := json.Marshal(book)
		handleError(e)
		modifiedData := append(data, encoded...)
		modifiedData = append(modifiedData, ']')
		file.Truncate(0)
		file.Seek(0, 0)
		file.Write(modifiedData)
	} else {
		jsonData, _ := json.Marshal([]p1.Book{book})
		file.Write(jsonData)
	}
	wg.Done()
}

func ReadFromJSON(ch chan p1.Book) {
	var books []p1.Book
	data, err := os.ReadFile(FILE_NAME)
	handleError(err)
	e := json.Unmarshal(data, &books)
	handleError(e)
	for i := 0; i < 5; i++ {
		ch <- books[i]
	}
}
