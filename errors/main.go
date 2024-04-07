package main

import (
	"fmt"
	"github.com/pkg/errors"
)

var errAnotherThing = errors.New("error doing another thing")

func doAnotherThing() error {
	return errAnotherThing
}

func doSomething() error {
	err := doAnotherThing()
	return fmt.Errorf("error doing something: %w", err)
}

func main() {
	err := doSomething()

	fmt.Println(err.Error())
	fmt.Println(errAnotherThing.Error())

	if errors.Is(err, errAnotherThing) {
		fmt.Println("Found error!")
	}

	if errors.Is(errAnotherThing, err) {
		fmt.Println("Found error 2!")
	}

	fmt.Println(err)
}
