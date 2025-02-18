/*
Named return values:
Sometimes, we can define the returning variable in the return type.
They are known as named return values.
But these statement should be only used in short function as it may raise the confusions.
ex:
*/

package main

import "fmt"

func calc(x, y, z int, a, b float64) (sum int, pro float64) {
	sum = x + y + z
	pro = a * b
	return
}

func main() {
	s, p := calc(5, 4, 3, 8, 8)
	fmt.Println("Sum: ", s, "\nproduct: ", p)
}
