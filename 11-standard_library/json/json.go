package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Order struct {
	ID          string    `json:"id"`
	DateOrdered time.Time `json:"date_ordered"`
	CustomerID  string    `json:"customer_id"`
	Items       []Item    `json:"items"`
}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	data := `{
		"id":"12345",
		"date_ordered":"2020-05-01T13:01:02Z",
		"customer_id":"3",
		"items":[{"id":"xyz123","name":"Thing 1"},{"id":"abc789","name":"Thing 2"}]
	}`

	var o Order
	err := json.Unmarshal([]byte(data), &o)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(o)

	out, err := json.Marshal(o)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
