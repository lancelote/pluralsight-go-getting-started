package main

func main() {
	// Old style
	mp := messagePrinter{"foo"}
	println(mp.message)

	// Using method
	mp2 := messagePrinter{"bar"}
	mp2.printMessage()
}

type messagePrinter struct {
	message string
}

func (mp *messagePrinter) printMessage() {
	println(mp.message)
}
