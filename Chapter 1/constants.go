/*
Constant variable value cannot be changed.

Syntax:
cont variableName = value

Point:
1. Name of constant variable should start with capital letter
2. constant cannot be declared by :=
3. scope of constant variable is limited to its block.
ex:
*/
package main

import "fmt"

const Area = 33

func main() {
	const Paramenter = 21
	fmt.Println("Current constant values: ", Area, Paramenter)
	// Area = 3 //Throw error
	// Paramenter = 2 //Throw error
}
