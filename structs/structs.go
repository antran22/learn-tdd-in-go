package structs

import (
	"math"
)

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2

}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return c.Radius * 2 * math.Pi
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

type RightTriangle struct {
	SideA float64
	SideB float64
}

func (r RightTriangle) Perimeter() float64 {
	hypotenuse := math.Sqrt(r.SideA*r.SideA + r.SideB*r.SideB)
	return r.SideA + r.SideB + hypotenuse
}

func (r RightTriangle) Area() float64 {
	return r.SideB * r.SideA / 2
}
