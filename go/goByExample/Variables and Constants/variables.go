package main

import "fmt"



func main() {
	/*
	Rules of naming Variables
	- A variable name consists of alphabets, digits, and an underscore
	- Variables cannot have other symbol ($,@,#,etc).
	- Variables name cannot begin with a number
	- A variable name cannot be a reserved word as they are part of the Go syntax like int,type,for,etc
	*/

	/*
	Data types 	Description 					Examples
	int			Integer numbers					7123, 0, -5, 7023
	float		Numbers and decimal points		20.2, 500.123456, -34.23
	complex		Complex numbers					2+4i, -9.5+18.3i
	string		Sequence of characters			"Hello World!"
												"1 is less than 2"
	bool		Either true or false			true, false
	*/
	
	var number int = 5
	var number2 = 10
	number3 := 15

	number2 = 20 // type cant be change

	fmt.Println(number)
	fmt.Println(number2)
	fmt.Println(number3)

	const Pi = 3.14	 // constant are fixed value that are not going to change
	// Pi = 4 // this cant be done, cannot assign to Pi (neither addressable nor a map index expression)
}

/*
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Variables\ and\ Constants/variables.go 
5
20
15
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Variables\ and\ Constants/variables.go 
# command-line-arguments
Variables and Constants/variables.go:26:2: cannot assign to Pi (neither addressable nor a map index expression)
*/