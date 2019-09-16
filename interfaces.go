package main

import (
    "fmt"
    "math"
)

// Geometry interface 
type Geometry interface {
    Area() float64 // - method signatures
    Perimeter() float64 // - method signatures
}

type Rect struct {
    width, height float64
}
type Circle struct {
    radius float64
}

// Interface method implementation
func (r Rect) Area() float64 {
    return r.width * r.height
}
func (r Rect) Perimeter() float64 {
    return 2*r.width + 2*r.height
}

// Interface method implementation
func (c Circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.radius
}


func measure(g Geometry) {
    fmt.Println(g.Area())
    fmt.Println(g.Perimeter())
}

func main() {
    r := Rect{width: 3.2, height: 4.5}
	c := Circle{radius: 5}

	measure(r)
    measure(c)
}