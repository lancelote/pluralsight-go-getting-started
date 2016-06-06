package main

import "fmt"

func main() {
	foo := myStruct{}
	foo.myField = "bar"
	fmt.Println("Struct foo:", foo)
	println("Struct field value:", foo.myField)

	foo2 := myStruct{"bar"}
	println("Struct2 field value:", foo2.myField)

	// Create struct on heap
	foo3 := new(myStruct)
	foo3.myField = "bar"
	println("Heap Struct field value:", foo3.myField)
	println("foo3 is a memory address:", foo3)

	// Another reference
	foo4 := &myStruct{"bar"}
	fmt.Println("Address reference:", foo4)
	fmt.Println("Adress reference struct field value:", foo4.myField)
}

type myStruct struct {
	myField string
}
