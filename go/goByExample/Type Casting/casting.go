package main

import "fmt"

func main() {
	var num1 float32 = 2.56
	var num2 int = int(num1)

	fmt.Printf("Int: %d\n", num2)

	var num3 int = 5
	var num4 float32 = float32(num3)

	fmt.Printf("Float: %f\n", num4)

}

/*
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Type\ Casting/casting.go
Int: 2
Float: 5.000000
*/
