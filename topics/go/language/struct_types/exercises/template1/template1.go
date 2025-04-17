// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

import "fmt"

// Add imports.

// Add user type and provide comment.
type user struct {
	name  string
	email string
	age   int
}

func main() {

	// Declare variable of type user and init using a struct literal.

	u1 := user{
		name:  "Bob",
		email: "bob@example.com",
		age:   20,
	}

	// Display the field values.
	fmt.Printf("u1: %+v\n", u1)

	// Declare a variable using an anonymous struct.
	u2 := struct {
		name  string
		email string
		age   int
	}{
		name:  "Dan",
		email: "dan@example.com",
		age:   21,
	}

	// Display the field values.
	fmt.Printf("u2: %+v\n", u2)
}
