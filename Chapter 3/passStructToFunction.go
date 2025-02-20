/*
Passing struct to function:
they can be passed by value(a copy of struct) or by reference using pointers

Pass by value--------------------------------------------
*/

// package main

// import "fmt"

// type student struct {
// 	Name   string
// 	rollno int
// 	age    int
// 	class  string
// }

// func printing(st student) {
// 	fmt.Println("name: ", st.Name)
// 	fmt.Println("Age: ", st.age)
// 	fmt.Println("Roll no: ", st.rollno)
// 	fmt.Println("Class: ", st.class)
// }

// func main() {
// 	s := student{Name: "Ravan", rollno: 12, age: 100, class: "X"}
// 	printing(s)
// }

/*
call by reference:........................................
pointer
ex:
*/

// package main

// import "fmt"

// type student struct {
// 	Name   string
// 	rollno int
// 	age    int
// 	class  string
// }

// func printing(st *student) { //Pointer storing the address of struct instance
// 	fmt.Println("name: ", st.Name)
// 	fmt.Println("Age: ", st.age)
// 	fmt.Println("Roll no: ", st.rollno)
// 	fmt.Println("Class: ", st.class)
// }

// func main() {
// 	s := student{Name: "Ravan", rollno: 12, age: 100, class: "X"}
// 	printing(&s) //Passing address of the struct instance
// }

/*
Anonymus struct:....................................................
Anonymus struct actually doesn't contains any name for the struct.,
means creating a struct without explicitly naming them
*/

// package main

// import "fmt"

// func main() {
// 	student := struct {
// 		Name   string
// 		age    int
// 		rollno int
// 		class  string
// 	}{
// 		Name:   "Ravan",
// 		age:    100,
// 		rollno: 12,
// 		class:  "X",
// 	}
// 	// student.Name = "Khali" // this line will change the value
// 	fmt.Println(student.Name)
// }

/*
Struct literals:
they denote a newly allocated struct value by listing the values of it fields.
You can list a subset of fields by using the Name: syntax
the special prefix & return a pointer to struct value

*/

package main

import "fmt"

type person struct {
	Name, Grp string
}

var (
	v1 = person{"Ravan", "Asur"}  // has type person
	v2 = person{Name: "Ravan"}    // Grp is empty
	v3 = person{Grp: "Asur"}      // Name is empty
	v4 = person{}                 //default
	p  = &person{"Ravan", "Asur"} //Has type *person
)

func main() {
	fmt.Println(v1, v2, v3, p)
}
