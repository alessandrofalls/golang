/*
Exercise 01

- Using the short declaration operator, assigns these values to variables with the identifiers "x", "y", and "z".
      1. 42
      2. "James Bond"
      3. true
- Now demonstrate the values in these variables, with:
      1. A single print statement
      2. Multiple print statements
*/
package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Println("Hello, Chico, the result of variables are:", x, y, z)
	fmt.Printf("The value of x is %v, and type of x is: %T\n", x, x)
	fmt.Printf("The value of y is %v, and type of y is: %T\n", y, y)
	fmt.Printf("The value of z is %v, and type of z is: %T\n", z, z)
}
