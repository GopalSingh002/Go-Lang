/*
maps in go lang, are data structure that allows you to store key value pair.

Syntax:..........................................................
var varname = map[key_type]value_type

initializing:....................................................
varname = make(map[typekey]typevalue)
	OR
varname :=map[keytype]valuetype{
key: value
key:value
...
key:value
}

Adding/Updating:.................................................
varname[key] = newvalue

Accessing Element:...............................................
in go lang, accessing map using key, return two things, 1st value, 2nd boolean specify whether the key exist or not

val, status :=m[key]
if status{
//found
}else{
//not found
}
print()
Deleting the pair:
-delete() function

delete(var, key)


*/

package main

import "fmt"

var mp1 = make(map[string]int)

func main() {
	mp := make(map[string]string)
	mp["Name"] = "Ravan"
	mp["Class"] = "12"
	mp["Age"] = "100"
	fmt.Println(mp)

	mp1["Kills"] = 1200
	mp1["Revive"] = 11
	mp1["Age"] = 50
	fmt.Println(mp1)

}
