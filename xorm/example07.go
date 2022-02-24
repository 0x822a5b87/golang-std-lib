package main

import (
	"fmt"
)

func example07() {
	rows, _ := engine.Where("age > ? and age < ?", 12, 30).Rows(&User{})
	defer rows.Close()

	u := &User{}
	for rows.Next() {
		rows.Scan(u)

		fmt.Println(u)
	}
}
