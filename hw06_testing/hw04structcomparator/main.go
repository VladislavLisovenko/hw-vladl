package hw04structcomparator

import (
	"errors"
)

type CompareMethod int

const (
	ByYear CompareMethod = 0
	BySize CompareMethod = 1
	ByRate CompareMethod = 2
)

type BookComparator struct {
	compareMethod CompareMethod
}

func NewBookComparator(cm CompareMethod) *BookComparator {
	return &BookComparator{
		compareMethod: cm,
	}
}

func (bc BookComparator) Compare(b1 *Book, b2 *Book) (bool, error) {
	switch bc.compareMethod {
	case ByYear:
		return b1.Year() > b2.Year(), nil
	case BySize:
		return b1.Size() > b2.Size(), nil
	case ByRate:
		return b1.Rate() > b2.Rate(), nil
	default:
		return false, errors.New("wrong compare method")
	}
}

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

func (b Book) ID() int {
	return b.id
}

func (b *Book) SetID(id int) {
	b.id = id
}

func (b Book) Title() string {
	return b.title
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b Book) Author() string {
	return b.author
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b Book) Year() int {
	return b.year
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b Book) Size() int {
	return b.size
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b Book) Rate() float64 {
	return b.rate
}

func (b *Book) SetRate(rate float64) {
	b.rate = rate
}
