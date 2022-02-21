package main

import (
	"log"
)

type User struct {
	Name string
	Age  int
}

func example01() {
	u := User{
		Name: "xxx",
		Age:  18,
	}

	log.Printf("%s login, age:%d", u.Name, u.Age)
	log.Panicf("Oh, system error when %s login", u.Name)
	log.Fatalf("Danger! hacker %s login", u.Name)
}
