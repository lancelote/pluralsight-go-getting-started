package main

func main() {
	var message = "Hello"

	sayHello(message)
	println(message)

	sayHelloByReference(&message)
	println("Mutated message:", message)

	variadicSayHello("Hello", "Variadic", "Function")
}

func sayHello(message string) {
	println(message)
	message = "Hello Go"
}

func sayHelloByReference(message *string) {
	println("Memory address:", message)
	*message = "Hello Go 2"
}

// Variadic function
func variadicSayHello(messages ...string) {
	for _, message := range messages {
		println(message)
	}
}
