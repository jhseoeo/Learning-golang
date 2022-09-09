package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	data := `
        { "name": "Fred", "age": 40 }
        { "name": "Mary", "age": 21 }
        { "name": "Pat", "age": 30 }
      `

	var t struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	dec := json.NewDecoder(strings.NewReader(data))
	for dec.More() {
		err := dec.Decode(&t)
		if err != nil {
			panic(err)
		}
		fmt.Println(t)
	}

	fmt.Println("---------------------------")

	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	var allInputs = []Person{
		{Name: "Fred", Age: 40},
		{Name: "Mary", Age: 21},
		{Name: "Pat", Age: 30},
	}

	process := func(a Person) Person {
		return a
	}

	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	for _, input := range allInputs {
		t := process(input)
		err := enc.Encode(t)
		if err != nil {
			panic(err)
		}
	}
	out := b.String()
	fmt.Println(out)
}
