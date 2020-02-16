package main

import (
	"fmt"

	"github.com/theskepticalowl/lib/readinput"
	"github.com/theskepticalowl/lib/simplemath"
)

func main() {
	//read float64 x for input
	fmt.Println("Please input X")
	var x float64 = readinput.ReadFloat()
	fmt.Println("")

	//read string calctype for operations
	fmt.Println("Type of calculation:")
	var calctype string = readinput.ReadString()
	fmt.Println("")

	//read float64 y for input
	fmt.Println("Please input Y")
	var y float64 = readinput.ReadFloat()
	fmt.Println("")

	//print output
	fmt.Println("Output:")

	if calctype == "+" {
		fmt.Println(simplemath.Add(x, y))
		return
	}
	if calctype == "-" {
		fmt.Println(simplemath.Subtract(x, y))
		return
	}
	if calctype == "*" {
		fmt.Println(simplemath.Multiply(x, y))
		return
	}

	if calctype == "/" {
		fmt.Println(simplemath.Divide(x, y))
		return
	}

	if calctype == "^" {
		fmt.Println(simplemath.Exponent(x, y))
		return
	}
}
