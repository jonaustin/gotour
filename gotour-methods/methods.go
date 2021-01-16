package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Vertex struct {
	X, Y float64
}

// Go doesn't have classes
// methods are just functions attached to types
// e.g. attach Abs() to Vertex var `v` - v.Abs()
// compare with normal func version:
//     func Abs(v Vertex) float64 {
func (v Vertex) Abs() float64 {
	sqrt := math.Sqrt(v.X*v.X + v.Y*v.Y)
	return sqrt
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type MyInt int

func (i MyInt) Abs() int {
	if i < 0 {
		return int(-i)
	}
	return int(i)
}

func main() {
	// struct
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	//float
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// int
	i := MyInt(rand.Int())
	fmt.Println(i.Abs())
}
