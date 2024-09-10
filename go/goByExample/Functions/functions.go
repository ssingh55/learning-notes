package main

import "fmt"

func greet() {
	// empty parentheses means no input is required to the function
	fmt.Println("Hello world!")
}

func addNumbers(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	var result int

	greet()
	result = addNumbers(20, 10)
	fmt.Println(result)
}

/*
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Functions/functions.go 
Hello world!
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Functions/functions.go 
Hello world!
30
*/