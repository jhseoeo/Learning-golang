package main

func main() {
	// using `type` keyword, we can define user-defined type
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}
	type Score int
	type Converter func(string) Score
	type TeamScores map[string]Score
	// we can only access the type from within its scope
}
