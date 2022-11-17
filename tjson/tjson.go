package main

import (
	"fmt"

	"hubble.com/json"
)

type Person struct {
	Name    string
	Age     int
	Grades  []int
	Parents []Person
}

func main() {
	p := Person{"Ivan", 22, []int{12, 15, 5}, []Person{}}
	jsonVal := json.ToJson(p)
	fmt.Println(jsonVal)
}
