package main

func main() {
	add := 1 + 2
	println("Addition:", add)

	subtract := 10 - 5
	println("Subtraction:", subtract)

	remainder := 5 % 2
	println("Remainder:", remainder)

	inc := 1
	inc++
	println("Increment 1:", inc)
	inc++
	println("Increment 2:", inc)
	// Cannot be used inline like `println(inc++)`

	// Augmented assignment
	inc += 5
	println("Augmented assignment:", inc)
}
