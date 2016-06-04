package main

func main() {
	result := add(1, 2, 3, 4, 5)
	println("Simple return:", result)

	numTerms, result := add2(6, 7, 8, 9, 3)
	println("Multiple return:", numTerms, "terms", result, "sum")

	numTerms, result = add3(3, 4, 5, 6, 7)
	println("Named return:", numTerms, "terms", result, "sum")
}

func add(terms ...int) int {
	var result = 0
	for _, term := range terms {
		result += term
	}
	return result
}

func add2(terms ...int) (int, int) {
	var result = 0
	for _, term := range terms {
		result += term
	}
	return len(terms), result
}

func add3(terms ...int) (numTerms int, sum int) {
	for _, term := range terms {
		sum += term
	}
	numTerms = len(terms)
	return
}
