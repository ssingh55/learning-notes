package main

import "fmt"


func main() {
	var score int
	fmt.Print("Enter a score: ")
	fmt.Scan(&score)

	var grade string
	switch {
	case score >= 90:
		grade = "A"
	case score >= 80:
		grade = "B"
	case score >= 70:
		grade = "C"
	case score >= 60:
		grade = "D"
	default:
		grade = "F"
	}

	fmt.Printf("Your grade is %s\n", grade)
}

/*
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Conditionals/switch.go 
Enter a score: 90
Your grade is A
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Conditionals/switch.go 
Enter a score: 88
Your grade is B
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Conditionals/switch.go 
Enter a score: 77
Your grade is C
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Conditionals/switch.go 
Enter a score: 65
Your grade is D
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Conditionals/switch.go 
Enter a score: 5
Your grade is F
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Conditionals/switch.go 
Enter a score: 33
Your grade is F
*/