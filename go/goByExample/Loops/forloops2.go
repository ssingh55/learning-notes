package main

import "fmt"

func main() {
	secretNumber := 27
	var guess int

	for {
		fmt.Print("Enter a guess: ")
		fmt.Scan(&guess)

		if guess == secretNumber {
			fmt.Println("Congratulations u guessed the secret number")
			break
		} else if guess < secretNumber {
			fmt.Println("Too low")
		} else {
			fmt.Println("too high")
		}
	}
}

/*
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Loops/forloops2.go 
Enter a guess: 34
too high
Enter a guess: 22
Too low
Enter a guess: 28
too high
Enter a guess: 27  
Congratulations u guessed the secret number
*/