package main

import "fmt"

func main() {
	// slices size once assigned can be changed
	numbers := []int{1, 2, 3}

	fmt.Println(numbers)

	numbers = append(numbers, 4, 5, 6)
	fmt.Println(numbers)
}

/*
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Arrays\ and\ Slices/slice.go 
[2 3 4]
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ 
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Arrays\ and\ Slices/slice.go 
[2 3 4]
[2 3 4 4 5 6]
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Arrays\ and\ Slices/slice.go 
[1 2 3]
[1 2 3 4 5 6]
*/