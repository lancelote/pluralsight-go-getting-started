package main

const (
	a = "the first"
	b = "the second"
	c = "the third"
)

const (
	first = 1 + iota
	second
	third
)

func main() {
	println(a)
	println(b)
	println(c)

	println(first)
	println(second)
	println(third)
}
