package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
)

func Example01Main() {
	type Option struct {
		Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug message"`
		Name    string `short:"n" long:"name" description:"your name"`
	}
	var opt Option
	_, err := flags.Parse(&opt)
	if err != nil {
		return
	}

	fmt.Println(opt.Name)
	fmt.Println(opt.Verbose)
}
