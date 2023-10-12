package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func function() error {
	return errors.New("function error")
}

func main() {
	err := function()
	if err != nil {
		err = errors.Wrapf(err, "wrap err")
		fmt.Printf("error = %+v\n", err)
	}
}
