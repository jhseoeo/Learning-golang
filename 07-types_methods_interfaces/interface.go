package main

import "fmt"

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	// do some business logic
	return data + "!"
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	// get data from somewhere
	data := "hello world"
	refinedData := c.L.Process(data)
	fmt.Println(refinedData)
}

func main() {
	c := Client{
		L: LogicProvider{}, // concrete type assigned into Client's interface
	}
	c.Program()
}
