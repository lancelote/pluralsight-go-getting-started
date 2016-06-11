package main

func main() {
	emp := enhancedMessagePrinter{}
	emp.message = "foo"

	// or
	// emp := enhancedMessagePrinter{messagePrinter{"foo"}}

	emp.printMessage()
}

type messagePrinter struct {
	message string
}

func (mp *messagePrinter) printMessage() {
	println(mp.message)
}

type enhancedMessagePrinter struct {
	messagePrinter
}
