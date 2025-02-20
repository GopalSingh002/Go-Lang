/*
Structs..........................................................
structs in go are the composite data type which group together different data type variables:

Syntax:
type structName struct{
	var1 datatype1
	var2 datatype2
}


-creating instance of struct:
var instancename structname

-initializing the struct variable:

instancename.var1 = value
instancename.var2 = value

				OR

instanceName:=structName{var1:value, var2:value}



-to access fields we can use the instance name
	instanceName.var

*/

package main

import "fmt"

type student struct {
	Name   string
	rollno int
	age    int
	class  string
}

func main() {
	var s student
	s.Name = "Ravan"
	s.age = 100
	s.rollno = 12
	s.class = "X"

	//accessing
	fmt.Println("name: ", s.Name)
	fmt.Println("Age: ", s.age)
	fmt.Println("Roll no: ", s.rollno)
	fmt.Println("Class: ", s.class)

}
