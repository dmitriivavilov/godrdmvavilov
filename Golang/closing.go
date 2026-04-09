package main

import "fmt"

func main() {
	fmt.Println("Hello")
	next := incrementor()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}

func incrementor() func() int {
	var i int
	fmt.Println("i before incrementing", i)
	return func() int {
		i++
		fmt.Println("i after incrementing", i)
		return i
	}
}
