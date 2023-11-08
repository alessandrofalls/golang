/*- Use var to declare three variables. They must have package-level scope. Do not assign values to these variables. Use the following identifiers and types for these variables:
     1. Identifier "x" must have type int
     2. Identifier "y" must have a string type
     3. Identifier "z" must have type bool
- In the main function:
     1. Demonstrate the values of each identifier
     2. The compiler assigned values to these variables. What are these values called?
*/
package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Printf("The value of x is: %v\n", x)
	fmt.Printf("The value of y is: %v\n", y)
	fmt.Printf("The value of z is: %v\n", z)
	fmt.Println("Hello, Chico")
}
