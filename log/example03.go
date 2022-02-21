package main

import (
	"bytes"
	"fmt"
	"log"
)

func example03() {
	u := User{
		Name: "dj",
		Age:  18,
	}
	buf := &bytes.Buffer{}
	logger := log.New(buf, "", log.Lshortfile|log.LstdFlags)

	logger.Printf("%s login, age:%d", u.Name, u.Age)

	fmt.Print(buf.String())
}
