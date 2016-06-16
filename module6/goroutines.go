package main

import (
	"time"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)

	go abcGen()

	println("This comes first!")

	time.Sleep(100*time.Millisecond)
}

func abcGen() {
	for l := byte('a'); l <= byte('z'); l++ {
		go println(string(l))
	}
}
