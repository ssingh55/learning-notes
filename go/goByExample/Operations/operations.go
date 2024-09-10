package main

import "fmt"

func main() {
	num1 := 20
	num2 := 10

	// Single line comment
	/*Double line comment*/

	// Arithemetic operations
	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)
	fmt.Println(num1 / num2)
	fmt.Println(num1 % num2)

	// logical oeprations
	fmt.Println(num1 == num2)
	fmt.Println(num1 != num2)
	fmt.Println(num1 > num2)
	fmt.Println(num1 < num2)
}

/*
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Operations/operations.go 
30
10
200
2
0
false
true
true
false
*/