package main

import "fmt"

func example05() {
	engine.Where("age > ? and age < ?", 12, 30).Iterate(&User{}, func(i int, bean interface{}) error {
		fmt.Println("user = ", i, ", value = ", bean.(*User))
		return nil
	})
}
