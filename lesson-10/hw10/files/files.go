package files

import (
	"encoding/json"
	"os"
)

type User struct {
	Id       uint
	Username string
	Fullname string
}
type Comment struct {
	Id     uint
	Body   string
	PostId uint
	Likes  uint
	User   User
}
type CommentsResponse struct {
	Comments []Comment
	Total    int
	Skip     int
	Limit    int
}

func JsonToStruct(filePath string) CommentsResponse {
	var comments CommentsResponse
	f, e := os.Open(filePath)
	if e != nil {
		panic(e)
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	err := decoder.Decode(&comments)
	if err != nil {
		panic(err)
	}
	return comments
}
