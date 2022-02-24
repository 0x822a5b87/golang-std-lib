package main

import "fmt"

func example03() {
	user1 := &User{}
	_, err := engine.ID(1).Cols("id", "name", "age").Get(user1)
	if err != nil {
		return
	}
	fmt.Printf("user1:%v\n", user1)

	user2 := &User{Name: "xxx"}
	_, err = engine.Omit("created", "updated").Get(user2)
	if err != nil {
		return
	}
	fmt.Printf("user2:%v\n", user2)
}
