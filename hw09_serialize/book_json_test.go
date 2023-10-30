package main

import (
	"reflect"
	"testing"

	"github.com/VladislavLisovenko/hw-vladl/hw09_serialize/bookjson"
	"github.com/stretchr/testify/require"
)

func MarshalUnmarshalJSONTests() []struct {
	descr string
	book  bookjson.Book
} {
	return []struct {
		descr string
		book  bookjson.Book
	}{
		{
			descr: "",
			book: bookjson.Book{
				ID:     1,
				Title:  "Title 1",
				Author: "Author 1",
				Year:   2020,
				Size:   202,
				Rate:   123.23,
			},
		},
		{
			descr: "",
			book: bookjson.Book{
				ID:     2,
				Title:  "Title 2",
				Author: "Author 2",
				Year:   2022,
				Size:   208,
				Rate:   323.27,
			},
		},
	}
}

func MarshalUnmarshalJSONSliceTests() []struct {
	desrc string
	slice []bookjson.Book
} {
	return []struct {
		desrc string
		slice []bookjson.Book
	}{
		{
			desrc: "",
			slice: []bookjson.Book{
				{
					ID:     1,
					Title:  "Title 1",
					Author: "Author 1",
					Year:   2020,
					Size:   202,
					Rate:   123.23,
				},
				{
					ID:     4,
					Title:  "Title 4",
					Author: "Author 4",
					Year:   2024,
					Size:   204,
					Rate:   123.24,
				},
			},
		},
	}
}

func TestMarshalUnmarshalJSON(t *testing.T) {
	tests := MarshalUnmarshalJSONTests()

	for _, td := range tests {
		td := td
		t.Run(td.descr, func(t *testing.T) {
			b, err := td.book.MarshalJSON()
			if err != nil {
				require.NoError(t, err)
				return
			}
			require.NotNil(t, b)

			var book bookjson.Book
			err = book.UnmarshalJSON(b)
			if err != nil {
				require.NoError(t, err)
				return
			}
			require.Equal(t, td.book, book)
		})
	}
}

func TestMarshalUnmarshalJSONSlice(t *testing.T) {
	tests := MarshalUnmarshalJSONSliceTests()

	for _, td := range tests {
		td := td
		t.Run(td.desrc, func(t *testing.T) {
			b, err := bookjson.MarshalSlice(td.slice)
			if err != nil {
				require.NoError(t, err)
				return
			}
			require.NotNil(t, b)

			slice, err := bookjson.UnmarshalSlice(b)
			if err != nil {
				require.NoError(t, err)
				return
			}
			for i := 0; i < len(slice); i++ {
				b1 := td.slice[i]
				b2 := slice[i]
				require.True(t, reflect.DeepEqual(b1, b2))
			}
		})
	}
}
