package main

import "fmt"

func main() {
	var num int

	fmt.Print("Enter the number: ")
	fmt.Scan(&num)

	// for initializingExpression; condition; increment/decrement {}
	for i := 1; i <= 10; i++ {
		if (i == 3) {
			continue
		}
		fmt.Printf("%d x %d = %d\n", num, i, num * i)
	}
}

/*
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Loops/forloops.go // Without continue
Enter the number: 2
2 * 1 = 2
2 * 2 = 4
2 * 3 = 6
2 * 4 = 8
2 * 5 = 10
2 * 6 = 12
2 * 7 = 14
2 * 8 = 16
2 * 9 = 18
2 * 10 = 20
shubham@shubham-prlappy:~/learning-notes/go/goByExample$ go run Loops/forloops.go 
Enter the number: 3
3 x 1 = 3
3 x 2 = 6
3 x 4 = 12
3 x 5 = 15
3 x 6 = 18
3 x 7 = 21
3 x 8 = 24
3 x 9 = 27
3 x 10 = 30

*/