package main

import (
	"fmt"
	"testing"

	sc "github.com/VladislavLisovenko/hw-vladl/hw04_struct_comparator"
)

func TestNewBookComparator(t *testing.T) {

	b1 := &sc.Book{}
	b1.SetID(1)
	b1.SetTitle("Book 1")
	b1.SetAuthor("Author 1")
	b1.SetYear(2020)
	b1.SetSize(10)
	b1.SetRate(1.1)

	b2 := &sc.Book{}
	b2.SetID(2)
	b2.SetTitle("Book 2")
	b2.SetAuthor("Author 2")
	b2.SetYear(2023)
	b2.SetSize(20)
	b2.SetRate(2.2)

	comparator := sc.NewBookComparator(sc.ByRate)
	r, err := comparator.Сompare(b1, b2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(r)

}
