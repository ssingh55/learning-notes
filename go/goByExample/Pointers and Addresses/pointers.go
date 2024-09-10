package main

import "fmt"

func main() {
	var number int = 5

	fmt.Println(number)
	fmt.Println(&number)

	var numPointer *int = &number

	fmt.Println("Using pointer: ", *numPointer)
	*numPointer = 10
	fmt.Println("Updated pointer values: ", *numPointer)
}


/*
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Pointers\ and\ Addresses/pointers.go 
5
0xc0000120a0
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Pointers\ and\ Addresses/pointers.go 
5
0xc0000120a0
0xc0000120a0
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Pointers\ and\ Addresses/pointers.go 
5
0xc0000120a0
Using pointer:  5
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Pointers\ and\ Addresses/pointers.go 
5
0xc00009c010
Using pointer:  5
Updated pointer values:  10
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Pointers\ and\ Addresses/pointers.go 
5
0xc0000120a0
Using pointer:  5
Updated pointer values:  0xc0000120a0
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Pointers\ and\ Addresses/pointers.go 
5
0xc0000120a0
Using pointer:  0xc0000120a0
Updated pointer values:  0xc0000120a0
*/