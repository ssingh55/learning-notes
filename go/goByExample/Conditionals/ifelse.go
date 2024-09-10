package main

import "fmt"

func main() {
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	// If (condition) {} else if (condition) {} else {}
	if (number > 0 ) {
		fmt.Println("Your number is positive")
	} else if (number < 0) {
		fmt.Println("Your number is negative")
	} else 	{
		fmt.Println("Your number is zero")
	}
}

/*
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Conditionals/ifelse.go 
Enter a number: 4
Your number is positive
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Conditionals/ifelse.go 
Enter a number: -5
Your number is negative
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Conditionals/ifelse.go 
Enter a number: 0
Your number is zero
*/