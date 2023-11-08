/*
- Using the solution from the previous exercise:
     1. In package-level scope, assign the following values to the variables:
         1. for "x" assign 42
         2. for "y" assign "James Bond"
         3. for "z" assign true
     2. In the main function:
         1. Use fmt.Sprintf to assign all these values to a single variable. Make this string type assignment to a variable named "s" using the short declaration operator.
         2. Demonstrate the variable "s".
*/
package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("The values of x, y and z in a row are: %v\t %v\t %v", x, y, z)
	fmt.Println(s)
}
