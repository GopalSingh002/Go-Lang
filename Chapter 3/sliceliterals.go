/*
Slice literals are like the array without a specific lenght
ex:
[]int{1,2,3,4,5}
*/
package main

import "fmt"

func main() {
	//Array
	q := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(q)
	r := []bool{true, false, false, true, true, true}
	fmt.Println(r)

	//accessing the struct elements
	person := struct {
		name string
		age  int
	}{
		name: "Ravan",
		age:  100,
	}

	fmt.Println(person)
	fmt.Println(person.name)

	//Accessing the array struct elements

	s := []struct {
		i int
		b bool
	}{
		{2, false},
		{3, true},
		{4, false},
		{5, true},
	}
	fmt.Println(s)
	fmt.Println(s[0])
	fmt.Println(s[0].i)
	fmt.Println(s[2].i)
	fmt.Println(s[0].b)

}
