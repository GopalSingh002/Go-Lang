/*
Go language support datatypes:
1. bool : true, false
2. string : "hello world"
3. int int8  int16 int32 int64
	int : it is platform dependent, usually 32, 64 bits
	intx, x(8,16,32,64) bit signed integers

4. uint uint8 uint16 uint32 uint64 uintptr
5. byte : allias of uint8
6. rune: allias of int32, represents in a unicode point
7. float32 float 64
8. complex64 complex128

*/
/*
Zero Values----------------------------------------------------
Variable declared without value still holds some default or initial values
which are called zero values
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	//booleans
	var a bool = true
	var b bool = false

	//string
	var c string = "hello world"

	//Integers (signed)
	var d int = 12 //platform dependent size
	var e int8 = -128
	var f int16 = 32000
	var g int32 = 100000
	var h int64 = 1000000000

	var x1 uint = 53 //platform dependant size
	var x2 uint8 = 255
	var x3 uint16 = 65535
	var x4 uint32 = 100000
	var x5 uint64 = 1000000000
	var x6 uintptr = 12312 //it is large enough to hold the bit pattern of any pointer

	var by byte = 65  // A in ascii
	var ru rune = 'A' // unicode of A

	var f1 float32 = 3.14
	var f2 float64 = math.Pi

	var comp1 complex64 = complex(3.0, 4.0)
	var comp2 complex128 = complex(3.0, 4.0)

	fmt.Println("Boolean: ", a, b)
	fmt.Println("String: ", c)
	fmt.Println("Integer(signed): ", d, e, f, g, h)
	fmt.Println("Unsigned: ", x1, x2, x3, x4, x5, x6)
	fmt.Println("Byte: ", by)
	fmt.Println("rune: ", ru)
	fmt.Println("float32: ", f1)
	fmt.Println("float64: ", f2)
	fmt.Println("complex64: ", comp1)
	fmt.Println("complex128: ", comp2)

	var i int
	var fl float32
	var s string
	var cm complex128
	var bl bool

	fmt.Println("\n\nZero values \n", "int: ", i, "\nFloat", fl, "\nString", s, "\nComplex: ", cm, "\nbool: ", bl)

}
