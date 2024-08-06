package main

import "fmt"

const (
	a = 2 << iota
	b
	c
	d
)

func main() {
	fmt.Println("hello btc")
	fmt.Println(b)
}
