package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}

// make a project as go module
// go mod init [projname]

// formatting code
// go fmt ./..
// goimports -l -w .

// linting
// golint ./.. (catches syntactic errors)
// go vet ./.. (catches wrong number of parameters, unused variables, etc..)
// golangci-lint run (both)
