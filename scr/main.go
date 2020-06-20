package main

import (
	"fmt"
	"strconv"
)

func main() {
	// in golang all variables MUST be used
	/*
		three typers of creating variable
			var i int
			var j int = 32
			k := 1
	*/
	k := 1 // 1
	// %v print variable
	// %T prints variable type
	fmt.Printf("%v, %T \n", k, k) // 1, int

	// to change int to float64 when
	// we use := operator just simply add .
	// after the value

	j := 11.                      // 11
	fmt.Printf("%v, %T \n", j, j) // 11, float64

	// also we can declare a group of variables like that
	/*
		var (
			m float32
			i int
		)
	*/

	// see 38th line
	fmt.Printf("%v, %T\n", a, a) // 10 int
	fmt.Printf("%v, %T\n", b, b) // 0, float64

	// converting variable types
	var a int = 10               // 10
	var b float32 = float32(a)   // 10
	fmt.Printf("%v, %T\n", b, b) // 10, float32

	// also if we have a float with value 4.5
	// the convertion to int will delete the value after . and just forget it
	var ass float32 = 4.5            // 4.5
	var bss int = int(ass)           // 4
	fmt.Printf("%v, %T\n", bss, bss) // 4, int

	// string convertion
	// if we have a integer and we want it to become a string
	// we need to convert it, but no simply doing string(integer)
	// thats because it takes integer unicode char
	var i = 42                   // 42, integer
	var l string = string(i)     // *, string
	fmt.Printf("%v, %T\n", l, l) // *, string

	// to convert it we need to import libriary srtconv
	l = strconv.Itoa(i)          // 42
	fmt.Printf("%v, %T\n", l, l) // 42, string

}

// also you can declare a variable outside the main function
// but you'll need to fully declare it. The usage if := operator
// will cause an error
var (
	a int = 10
	b float64
)

// by default b = 0

// lowercase variable means it is seen only in the same package
// while uppercase means its a global one
