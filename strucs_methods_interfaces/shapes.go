package shapes

import "math"

// Shape interface
type Shape interface {
	Area() float64
}

// Rectangle struct
type Rectangle struct {
	Width float64
	Height float64
}

// Circle struct
type Circle struct {
	Radius float64
}

// Triangle struct
type Triangle struct {
	Base float64
	Height float64
}

// Perimeter calc
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area rectangle calc
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area circle calc
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area triangel calc
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}