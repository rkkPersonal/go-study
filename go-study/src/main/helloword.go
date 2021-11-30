package main

import "fmt"

var (
	firstName, lastName, age = "John", "Doe", 32
)

func main() {
	fmt.Printf("hello, world\n")
	fmt.Printf(firstName, lastName, age)
}