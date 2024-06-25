package main

import "fmt"

func helloMain() {
	fmt.Println("Hello World!")
	fmt.Println(hello("amit"))
}

func hello(name string) string {
	message := fmt.Sprintf("hello, %v. Welcome to the Go!", name)
	return message
}
