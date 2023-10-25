package main

import (
	"testing"

	"github.com/VladislavLisovenko/hw-vladl/hw09_serialize/model/model"
	"github.com/stretchr/testify/require"
)

func MarshalUnmarshalJSONTests() []struct {
	descr string
	book  Book
} {
	return []struct {
		descr string
		book  Book
	}{
		{
			descr: "",
			book: Book{
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
			book: Book{
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

func MarshalUnmarshalSliceTests() []struct {
	desrc string
	slice []*model.Book
} {
	return []struct {
		desrc string
		slice []*model.Book
	}{
		{
			desrc: "",
			slice: []*model.Book{
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

			var book Book
			err = book.UnmarshalJSON(b)
			if err != nil {
				require.NoError(t, err)
				return
			}
			require.Equal(t, td.book, book)
		})
	}
}

func TestMarshalUnmarshalSlice(t *testing.T) {
	tests := MarshalUnmarshalSliceTests()

	for _, td := range tests {
		td := td
		t.Run(td.desrc, func(t *testing.T) {
			b, err := MarshalSlice(td.slice)
			if err != nil {
				require.NoError(t, err)
				return
			}
			require.NotNil(t, b)

			slice, err := UnmarshalSlice(b)
			if err != nil {
				require.NoError(t, err)
				return
			}
			for i := 0; i < len(slice); i++ {
				b1 := td.slice[i]
				b2 := slice[i]
				require.Equal(t, b1.ID, b2.ID)
				require.Equal(t, b1.Title, b2.Title)
				require.Equal(t, b1.Author, b2.Author)
				require.Equal(t, b1.Year, b2.Year)
				require.Equal(t, b1.Size, b2.Size)
				require.Equal(t, b1.Rate, b2.Rate)
			}
		})
	}
}
