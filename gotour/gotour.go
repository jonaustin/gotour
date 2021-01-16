package main

import (
	"fmt"
	"math"
	"strings"
	//"os"
)

var c, python, java = true, false, "no!"
var i, j int = 1, 2

//var i, j = 1, 2 # same thing; implicity typing

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return // implicitly returns named return args x,y
}

func loop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}

func whileLoop() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}

func infiniteLoop() {
	for {
	}
}

func ifFunc(x float64) string {
	if x < 0 {
		return ifFunc(-x) + "str"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func ifWithInit(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		// Variables declared inside an if short statement are also available inside any of the else blocks.
		fmt.Printf("%g >= %g\n", v, limit)
	}
	// can't use v here, though
	return limit
}

func deferred() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

func pointers() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

// structs
type Vertex struct {
	X int
	Y int
}

func structLiterals() {
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p  = &Vertex{1, 2} // has type *Vertex
	)

	fmt.Println(v1, p, v2, v3)
}

func array() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slice() {
	// An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.
	// In practice, slices are much more common than arrays.

	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	// Other slices that share the same underlying array will see those changes.

	primes := [6]int{2, 3, 5, 7, 11, 13} // array literal

	var s []int = primes[1:4] // slice

	fmt.Println(s)

	s2 := s[:2]
	fmt.Println(s2)

	s3 := s[1:]
	fmt.Println(s3)

	var sl = []int{1, 2, 3, 4} // slice literal
	fmt.Println(sl)

	// nil slice
	var slnil []int
	fmt.Println(slnil, len(slnil), cap(slnil))
	if slnil == nil {
		fmt.Println("nil!")
	}

	printSlice(s)
}

func printSlice(s []int) {
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func makeSlice() {
	// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
	// The make function allocates a zeroed array and returns a slice that refers to that array:
	a := make([]int, 5) // len(a)=5
	printSlice2("a", a)

	//specify capacity with 3rd arg to make
	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func slicesofSlicesTicTacToe() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func sliceAppend() {
	var s []int
	printSlice(s)

	// append works on nil slices
	s = append(s, 0)
	printSlice(s)

	// slice grows as needed
	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)
}

func sliceRange() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// _
	pow = make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func hashmap() {
	// A map maps keys to values.
	//The zero value of a map is nil. A nil map has no keys, nor can keys be added.
	//The make function returns a map of the given type, initialized and ready for use.
	type Vertex struct {
		Lat, Long float64
	}

	var m map[string]Vertex // map where keys are string and values are Vertex

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])
}

func mapLiteral() {
	// Map literals are like struct literals, but the keys are required.
	type Vertex struct {
		Lat, Long float64
	}

	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}

	fmt.Println(m)

	// If the top-level type is just a type name, you can omit it from the elements of the literal.
	m = map[string]Vertex{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)

	// mutation
	m2 := make(map[string]int)

	m2["answer"] = 42
	fmt.Println("The value:", m2["answer"])

	m2["answer"] = 48
	fmt.Println("The value:", m2["answer"])

	delete(m2, "answer")
	fmt.Println("The value:", m2["answer"])

	// Test that a key is present with a two-value assignment:
	v, isPresent := m2["answer"]
	fmt.Println("The value:", v, "Present?", isPresent)
}

// Functions are values too. They can be passed around just like other values.
// Function values may be used as function arguments and return values.
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func higherOrderFunctions() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

// closures
//  Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
// For example, the adder function returns a closure. Each closure is bound to its own sum variable.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func closures() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func main() {
	//fmt.Println(add(1, 2))

	//a, b := swap("hello", "world")
	//fmt.Println(a, b)

	//fmt.Println(split(9))

	fmt.Println(math.Pi)

	var i int
	k := 3 // duck (structural cuz at compile time) typing; only works inside a function

	fmt.Println(i, j, k, c, python, java)

	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// Unlike in C, in Go assignment between items of different type requires an explicit conversion
	var x, y int = 3, 4
	var fl float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(fl)
	fmt.Println(x, y, z, fl)

	v := "42" // change me!
	fmt.Printf("v is of type %T\n", v)

	const Pi = 3.14
	fmt.Println(Pi)

	loop()
	whileLoop()

	fmt.Println(ifFunc(2), ifFunc(-4))

	fmt.Println(
		ifWithInit(3, 2, 10),
		ifWithInit(3, 3, 20),
	)

	deferred()

	pointers()

	// structs
	fmt.Println(Vertex{1, 2})
	ver := Vertex{1, 2}
	ver.X = 4
	fmt.Println(ver.X)

	// struct pointers
	ver2 := Vertex{1, 2}
	p := &ver2
	p.X = 1e9
	fmt.Println(ver2)

	structLiterals()

	array()
	slice()
	makeSlice()
	slicesofSlicesTicTacToe()
	sliceAppend()
	sliceRange()

	hashmap()
	mapLiteral()

	higherOrderFunctions()
}
