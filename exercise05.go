/*
- Using the solution from the previous exercise:
 1. In package-level scope, using the var keyword, create a variable with the identifier "y". The type of this variable must be the underlying type of the type you created in the previous exercise.
 2. In the main function:
 1. This should already be done:
 1. Demonstrate the value of the variable "x"
 2. Demonstrate the type of variable "x"
 3. Assign 42 to the variable "x" using the "=" operator
 4. Demonstrate the value of the variable "x"
 2. Now also do:
 1. Use conversion to transform the type of the value of the variable "x" into its underlying type and, using the "=" operator, assign the value of "x" to "y"
 2. Demonstrate the value of "y"
 3. Demonstrate the type of "y"
*/
package main

import "fmt"

type myinteger int

var x myinteger
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("The type of var x is: %T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println("the value of y is:", y)
	fmt.Printf("And the type of y is %T", y)
}
