package bookproto

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

func MarshalSlice(books []*Book) ([]byte, error) {
	bookList := BookList{
		Books: books,
	}
	return proto.Marshal(&bookList)
}

func UnmarshalSlice(b []byte) ([]*Book, error) {
	var bookList = &BookList{}
	err := proto.Unmarshal(b, bookList)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return bookList.Books, nil
}
