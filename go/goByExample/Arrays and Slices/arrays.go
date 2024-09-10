package main

import "fmt"

func main() {
	// var name_of_arr = [size of array]datatype{v1, v2, v3, ......}
	// array size once assigned can't be changed
	var arrayOfAges = [3]int{20, 22, 24}

	fmt.Println(arrayOfAges)
	fmt.Println(arrayOfAges[1])

	arrayOfAges[2] = 26
	fmt.Println(arrayOfAges)

	var arrayOfScores [3]int
	arrayOfScores[0] = 98

	fmt.Println(arrayOfScores)

	for i := 0; i < len(arrayOfScores); i++ {
		fmt.Println(arrayOfAges[i])
	}
}

/*
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Arrays\ and\ Slices/arrays.go 
[20 22 24]
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Arrays\ and\ Slices/arrays.go 
[20 22 24]
22
26
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Arrays\ and\ Slices/arrays.go 
[20 22 24]
22
[20 22 26]
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Arrays\ and\ Slices/arrays.go 
[20 22 24]
22
[20 22 26]
[98 0 0]
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Arrays\ and\ Slices/arrays.go 
[20 22 24]
22
[20 22 26]
[98 0 0]
[20 22 26]
[20 22 26]
[20 22 26]
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Arrays\ and\ Slices/arrays.go 
[20 22 24]
22
[20 22 26]
[98 0 0]
20
22
26
*/