package main

import (
	"errors"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) Jump() {
	fmt.Println(p.Name, "is jumping.")
}

func (p Person) PickUpBeer() (bool, error) {
	if p.Age < 21 {
		return false, errors.New("My fake id is not working.")
	}
	return true, nil
}

func types() {
	fmt.Println("This is user defined types.")
	p := Person{
		Name: "John",
		Age:  30,
	}

	p.Jump()
	fmt.Println(p.PickUpBeer())
}
