/*
Multiple results:
A function can return multiple result of simillar and multiple types:
ex:
*/

package main

import "fmt"

func calc(x, y, z int, a, b float64) (int, float64) {
	return x + y + x, a * b
}
func main() {
	sum, pro := calc(2, 3, 4, 5.3, 9.2)
	fmt.Println("Sum: ", sum, "\nProduct: ", pro)
}
