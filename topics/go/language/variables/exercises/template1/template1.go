// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

import "fmt"

// Add imports

// main is the entry point for the application.
func main() {

	// Declare variables that are set to their zero value.
	var a string
	var b int
	var c bool

	// Display the value of those variables.
	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	aa := "Hello"
	bb := 10
	cc := true

	// Display the value of those variables.
	fmt.Printf("aa := \"Hello\" \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := 10 \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := true \t %T [%v]\n", cc, cc)

	// Perform a type conversion.
	dd := float64(bb)

	// Display the value of that variable.
	fmt.Printf("dd := float64(bb) \t %T [%v]\n", dd, dd)
}
