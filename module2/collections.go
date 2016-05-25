package main

import "fmt"

func main() {
	// Array consists of elements of the same type
	// Has fixed size and not commonly used
	myArray := [3]int{}
	myArray[0] = 42
	myArray[1] = 27
	myArray[2] = 99
	fmt.Println("myArray:", myArray)

	mySecondArray := [...]int{1, 2, 3}
	fmt.Println("mySecondArray:", mySecondArray)

	// Array length
	fmt.Println("myArray length:", len(myArray))

	// Slices - array subset
	mySlice := myArray[:]
	fmt.Println("Slice:", mySlice)

	// Add elements to slice
	mySlice = append(mySlice, 100)
	fmt.Println("myArray after append:", myArray)
	fmt.Println("mySlice after append:", mySlice)

	// Create slice form scratch
	mySecondSlice := []float32{14., 15., 19.}
	fmt.Println("mySecondSlice:", mySecondSlice)
	fmt.Println("mySecondSlice length:", len(mySecondSlice))

	// Create slice with given number of items predefined
	myThirdSlice := make([]float32, 3, 100)
	myThirdSlice[0] = 11.
	myThirdSlice[1] = 12.
	myThirdSlice[2] = 13.
	fmt.Println("myThirdSlice:", myThirdSlice)

	// Maps
	myMap := make(map[int]string)
	myMap[42] = "Foo"
	myMap[12] = "Bar"
	fmt.Println("myMap:", myMap)

	// Provides placeholder for unknown key
	fmt.Println("Placeholder:", myMap[99])
}
