// Stringers
//
// One of the most ubiquitous interfaces is Stringer defined by the fmt package.
//
// type Stringer interface {
//     String() string
// }
//
// A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

//      arg, type-to-add-method-to    method-name, return type
//func (p    Person)                  String()     string {
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type ErrNegativeInt float64

func (e ErrNegativeInt) Error() string {
	msg := fmt.Sprint("Age cannot be a negative number: ", int(e))
	return msg
}

// errors - https://tour.golang.org/methods/20
func (p Person) StringWithError() string {
	if p.Age < 0 {
		return ErrNegativeInt(p.Age).Error()
	}

	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// Exercise https://tour.golang.org/methods/18

type IPAddr [4]byte

func (ip IPAddr) String() string {
	str_slice := make([]string, len(ip))

	for i, v := range string(ip[:]) {
		str := strconv.Itoa(int(v))
		str_slice[i] = str
	}

	return strings.Join(str_slice, ".")
}

// EX

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	ip := IPAddr{127, 0, 0, 1}
	fmt.Println(ip.String())
	ip = IPAddr{1, 2, 3, 4}
	fmt.Println(ip.String())

	negAge := Person{"Tenet", -42}
	fmt.Println(negAge.StringWithError())
}
