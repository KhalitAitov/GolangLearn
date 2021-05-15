package main

import "fmt"

func main() {
	var a int
	var b int
	var op string
	//op ="/" "*" "-" "+"
	op = "/"
	a = 333
	b = 4

	if op == "-" {
		fmt.Println(a - b)
	} else if op == "+" {
		fmt.Println(a + b)
	} else if op == "*" {
		fmt.Println(a * b)
	} else if op == "/" {
		fmt.Println(a / b)
	}

}
