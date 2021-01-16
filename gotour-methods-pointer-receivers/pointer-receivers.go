// https://tour.golang.org/methods/7
//Functions that take a value argument must take a value of that specific type:

//var v Vertex
//fmt.Println(AbsFunc(v))  // OK
//fmt.Println(AbsFunc(&v)) // Compile error!

//while methods with value receivers take either a value or a pointer as the receiver when they are called:

//var v Vertex
//fmt.Println(v.Abs()) // OK
//p := &v
//fmt.Println(p.Abs()) // OK

// https://tour.golang.org/methods/6
package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p, *p)
}

/*
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Use a Pointer-Receiver when you need to _modify_ a value
// Methods with pointer receivers can modify the value to which the receiver points (as Scale does here).
// Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
// Note: methods with pointer receivers take either a value or a pointer as the receiver when they are called
//       (in contrast to function version below)
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// does nothing without *
// With a value receiver, the Scale method operates on a copy of the original Vertex value.
// (This is the same behavior as for any other function argument.)
// The Scale method must have a pointer receiver to change the Vertex value declared in the main function.
func (v Vertex) NoScale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.X, v.Y) // (3*3) + (4*4) => 25 => sqrt => 5
	fmt.Println(v.Abs())

	v.NoScale(10)
	fmt.Println(v.X, v.Y)
	fmt.Println(v.Abs())

	v.Scale(10)
	fmt.Println(v.X, v.Y)
	fmt.Println(v.Abs()) // (3*10 * 3*10) + (4*10 * 4*10) => 2500 => sqrt => 50
}
*/

// Function version
/*
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}

	// functions with a pointer argument must take a pointer:
	Scale(&v, 10) // note: this int is implicitly converted to float64

	fmt.Println(Abs(v))
}
*/
