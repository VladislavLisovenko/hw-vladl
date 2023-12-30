package main

import (
	"testing"

	"github.com/VladislavLisovenko/hw-vladl/hw09_serialize/bookproto"
	"github.com/stretchr/testify/require"
)

func MarshalUnmarshalSliceTests() []struct {
	desrc string
	slice []*bookproto.Book
} {
	return []struct {
		desrc string
		slice []*bookproto.Book
	}{
		{
			desrc: "",
			slice: []*bookproto.Book{
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

func TestMarshalUnmarshalSlice(t *testing.T) {
	tests := MarshalUnmarshalSliceTests()

	for _, td := range tests {
		td := td
		t.Run(td.desrc, func(t *testing.T) {
			b, err := bookproto.MarshalSlice(td.slice)
			if err != nil {
				require.NoError(t, err)
				return
			}
			require.NotNil(t, b)

			slice, err := bookproto.UnmarshalSlice(b)
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
