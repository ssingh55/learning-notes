package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}

/*
shubham@shubham-prlappy:~/learning-notes/go/goTheoryAndExample/project$ go build 
shubham@shubham-prlappy:~/learning-notes/go/goTheoryAndExample/project$ ls
go.mod  main.go  project
shubham@shubham-prlappy:~/learning-notes/go/goTheoryAndExample/project$ 
shubham@shubham-prlappy:~/learning-notes/go/goTheoryAndExample/project$ ./project 
Hello World!
shubham@shubham-prlappy:~/learning-notes/go/goTheoryAndExample/project$ go build -o mybinary
shubham@shubham-prlappy:~/learning-notes/go/goTheoryAndExample/project$ ls
go.mod  main.go  mybinary  project
shubham@shubham-prlappy:~/learning-notes/go/goTheoryAndExample/project$ ./mybinary 
Hello World!

shubham@shubham-prlappy:~/learning-notes/go/goTheoryAndExample/project$ GOOS=linux GOARCH=arm64 go build
shubham@shubham-prlappy:~/learning-notes/go/goTheoryAndExample/project$ ./project 
bash: ./project: cannot execute binary file: Exec format error
*/