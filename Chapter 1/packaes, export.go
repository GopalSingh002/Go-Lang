/*
Packages, Variables, Functions
*/

/*Packages---------------------------------------------------------
1 Every Go program is made up of Packages
2 program starts running in package main
3. By convensions, the package name is the same as the last element of the import path.
ex: math/rand = package:package rand
*/
/*
Exported names-------------------------------------------------------
In go, a name exported if it begins with capital letter.
Pi in math package.
ex:
*/

package main
import {
	"fmt"
	"mat"
}
func main(){
	fmt.Println(math.Pi)
}


/*
Elaboration: 

package main: This declare that the code is the part of main package , main package is the special because it contains the entry point of the program(main function),
like when you execute the program it execute from the main function

import {"fmt" "math"}: import statment allow us to use external packages in our program,
like here we used fmt, math packages.
fmt:for formatted I/O operation(Println, Printf)
math: Contains mathematical constants and function. like we used math.Pi

func main(){}:  this is the main function from where the program will starts to execute.
fmt.Println(): it print the value to the console
math.Pi(): this function is used to get the value of pi: 3.14.....

*/



