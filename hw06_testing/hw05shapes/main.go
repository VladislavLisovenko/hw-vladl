package hw05shapes

import (
	"errors"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() float64 {
	return (t.base * t.height) / 2
}

func CalculateArea(s any) (float64, error) {
	switch o := s.(type) {
	case Shape:
		return o.Area(), nil
	default:
		return 0.0, errors.New("it's not a Shape")
	}
}

type Square struct {
	side float64
}
