package bookjson

import "encoding/json"

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

func MarshalSlice(books []Book) ([]byte, error) {
	return json.Marshal(books)
}

func UnmarshalSlice(b []byte) ([]Book, error) {
	var books []Book
	err := json.Unmarshal(b, &books)
	if err != nil {
		return nil, err
	}

	return books, nil
}
