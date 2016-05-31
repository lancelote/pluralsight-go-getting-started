package main

func main() {
	// If statement
	if foo := 2; foo == 1 {
		println("True")
	} else {
		println("False")
	}

	// Switch statement
	switch foo := 1; foo {
	case 1:
		println("Foo is one")
	case 2:
		println("Foo is two")
	}

	var bar = 3
	switch {
	case bar == 1:
		println("Bar is one")
	case bar > 2:
		println("Bar is greater than 2")
	}
}
