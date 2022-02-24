package main

import (
	"fmt"
)

func example06() {
	age := 10
	num, _ := engine.Where("age >= ?", age).Count(&User{})
	logger.Info(fmt.Sprint("there are ", num, " users whose age >=", age))
}
