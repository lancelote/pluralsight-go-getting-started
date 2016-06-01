package main

func main() {
	println("For loop")
	for i := 0; i < 5; i++ { // Elements can be ommited
		println(i)
	}

	println("\nInfinit for loop aka while")
	var i = 0
	for {
		i++
		println(i)
		if i > 5 {
			break
		}
	}

	println("\nIterating over slice")
	var s = []string{"foo", "bar", "buz"}
	for idx, v := range s {
		println(idx, v)
	}

	println("\nIterate over map")
	var m = make(map[string]string)
	m["first"] = "foo"
	m["second"] = "bar"
	m["third"] = "buz"
	for k, v := range m {
		println(k, v)
	}
}
