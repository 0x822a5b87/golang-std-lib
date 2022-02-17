package main

import (
	"flag"
	"fmt"
)

func example04Main() {
	args := []string{"-intFlag", "12", "-stringFlag", "test"}

	var intFlag int
	var boolFlag bool
	var stringFlag string

	fs := flag.NewFlagSet("MyFlagSet", flag.ContinueOnError)
	fs.IntVar(&intFlag, "intFlag", 0, "int flag value")
	fs.BoolVar(&boolFlag, "boolFlag", false, "bool flag value")
	fs.StringVar(&stringFlag, "stringFlag", "default", "string flag value")

	err := fs.Parse(args)
	if err != nil {
		return
	}

	fmt.Println("int flag:", intervalFlag)
	fmt.Println("bool flag:", boolFlag)
	fmt.Println("string flag:", stringFlag)
}
