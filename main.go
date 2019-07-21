package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Age  int `json:"sAge"`
}

func main() {
	firstPerson := person{"Name01", 20}
	bs, _ := json.Marshal(firstPerson)

	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))
}
