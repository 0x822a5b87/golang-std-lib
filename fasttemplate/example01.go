package main

import (
	"fmt"
	"github.com/valyala/fasttemplate"
)

func example01() {
	template := `name: {{name}}
age: {{age}}`
	t := fasttemplate.New(template, "{{", "}}")
	s1 := t.ExecuteString(map[string]interface{}{
		"name": "dj",
		"age":  "18",
	})
	s2 := t.ExecuteString(map[string]interface{}{
		"name": "hjw",
		"age":  "20",
	})
	fmt.Println("==========")
	fmt.Println(s1)
	fmt.Println("==========")
	fmt.Println(s2)
	fmt.Println("==========")
}
