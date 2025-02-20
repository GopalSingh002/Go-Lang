/*
Pointer..............................................................
A pointer is a variable in Go lang, who store the memory address

Syntax:
-var variable_name *datatype
	ex: var p *int: pointer p store nil value till now, but it is assigned to int datatype. means it can store the address of int variable

- pvar = &var2
	pvar(pointer variable) store the address of var2. Now pvar is no longer nil

ex: var i int
	p = &i

- *p: this is called dereferencing operation on the pointer, means we are accessing the value stored at the address of pointer

- *p = 12
	now, here we are modifying the value of i using the pointer.
ex:
*/

package main

import "fmt"

func main() {
	var i int
	var p *int
	fmt.Println("Value of i before pointer assigning: ", i)
	p = &i
	fmt.Println("Pointer p, have the address of i now")
	fmt.Println("Current value of pointer p: ", *p, "Address stored by p: ", p)
	*p = 12
	fmt.Println("i value after dereferencing: ", i)
	fmt.Println("p value after dereferencing: ", *p, "Address stored by p: ", p)
}
