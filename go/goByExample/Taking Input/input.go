package main

import "fmt"

func main() {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	fmt.Printf("You are %d years old", age)
}

/*
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Taking\ Input/input.go 
Taking Input/input.go:3:11: missing import path
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Taking\ Input/input.go 
Enter your age: 23
You are 23 years old
*/