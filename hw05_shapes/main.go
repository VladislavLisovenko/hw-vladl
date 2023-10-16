package main

import (
	"errors"
	"fmt"
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

func calculateArea(s any) (float64, error) {
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

func main() {
	circle := &Circle{
		radius: 5,
	}
	area, err := calculateArea(circle)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Circle: radius %f\nArea %f\n\n", circle.radius, area)

	rect := &Rectangle{
		width:  10,
		height: 5,
	}
	area, err = calculateArea(rect)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Rectangle: width %f, height %f\nArea %f\n\n", rect.width, rect.height, area)

	triangle := &Triangle{
		base:   8,
		height: 6,
	}
	area, err = calculateArea(triangle)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Triangle: base %f, height %f\nArea %f\n\n", triangle.base, triangle.height, area)

	square := &Square{
		side: 6,
	}
	area, err = calculateArea(square)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Square: side %f\nArea %f", square.side, area)
}
