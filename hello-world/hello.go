package main

import "fmt"

func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	x := "Huy"
	fmt.Println(Hello(x))
}
