package main

import (
	"encoding/json"
	"fmt"

	"github.com/VladislavLisovenko/hw-vladl/hw09_serialize/model/model"
	"google.golang.org/protobuf/proto"
)

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Size   int     `json:"size"`
	Rate   float64 `json:"rate"`
}

func (book Book) MarshalJSON() ([]byte, error) {
	type BookTmp Book
	bookTmp := BookTmp(book)
	return json.Marshal(bookTmp)
}

func (book *Book) UnmarshalJSON(data []byte) error {
	type BookTmp Book
	var bookTmp BookTmp
	err := json.Unmarshal(data, &bookTmp)
	if err != nil {
		return err
	}

	*book = Book(bookTmp)
	return nil
}

func MarshalSlice(books []*model.Book) ([]byte, error) {
	bookList := model.BookList{
		Books: books,
	}
	return proto.Marshal(&bookList)
}

func UnmarshalSlice(b []byte) ([]*model.Book, error) {
	var bookList = &model.BookList{}
	err := proto.Unmarshal(b, bookList)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return bookList.Books, nil
}

func main() {

	book := Book{
		ID:     1,
		Title:  "book title",
		Author: "book author",
		Year:   2000,
		Size:   50,
		Rate:   45.123,
	}

	b, err := book.MarshalJSON()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	var book2 Book
	err = book2.UnmarshalJSON(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(book2)

	bookList := []*model.Book{
		{
			ID:     11,
			Title:  "title 11",
			Author: "author 11",
			Year:   2018,
			Size:   400,
			Rate:   3.27,
		},
		{
			ID:     15,
			Title:  "title 15",
			Author: "author 15",
			Year:   2015,
			Size:   500,
			Rate:   5.25,
		},
	}

	b2, err := MarshalSlice(bookList)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b2)

	bl, err := UnmarshalSlice(b2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bl)
}
