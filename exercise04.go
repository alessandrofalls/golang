/*
- Create a type. The underlying type must be int.
- Create a variable for this type, with the identifier "x", using the var keyword.
- In the main function:
     1. Demonstrate the value of the variable "x"
     2. Demonstrate the type of variable "x"
     3. Assign 42 to the variable "x" using the "=" operator
     4. Demonstrate the value of the variable "x"
*/
package main

import "fmt"

type myinteger int
var x myinteger

func main() {
	fmt.Println(x)
	fmt.Printf("The type of var x is: %T\n", x)
	x = 42
	fmt.Println(x)
}
