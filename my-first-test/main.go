package main

import "fmt"

const (
	first = iota
	second
	thread
	fourth
)

func main() {
	fmt.Println(2 << first)
	fmt.Println(2 << second)
	fmt.Println(2 << thread)
	fmt.Println(2 << fourth)
}
