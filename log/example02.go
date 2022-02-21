package main

import "log"

func example02() {
	u := User{
		Name: "dj",
		Age:  18,
	}

	log.SetPrefix("Login: ")
	log.SetFlags(log.Lshortfile)

	log.Printf("%s login, age:%d", u.Name, u.Age)
}
