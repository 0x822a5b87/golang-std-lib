package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"log"
)

func example03() {
	type Option struct {
		Required string `short:"r" long:"required" required:"true"`
		Default  string `short:"d" long:"default" default:"default"`
	}

	var opt Option
	_, err := flags.Parse(&opt)
	if err != nil {
		log.Fatal("Parse error:", err)
	}

	fmt.Println("required: ", opt.Required)
	fmt.Println("default: ", opt.Default)
}
