package main

import (
	"fmt"
	"time"
)

func example02() {
	user1 := &User{}

	// 使用主键ID搜索
	has, _ := engine.ID(1).Get(user1)
	if has {
		fmt.Printf("user1:%v\n", user1)
	} else {
		one, err := engine.InsertOne(User{
			Id:      1,
			Name:    "xxx",
			Salt:    "xxx",
			Age:     20,
			Passwd:  "123456",
			Created: time.Time{},
			Updated: time.Time{},
		})
		if err != nil {
			return
		}
		logger.Info(fmt.Sprint("one : ", one))
	}

	// 使用 where 条件搜索
	user2 := &User{}
	has, _ = engine.Where("name=?", "xxx").Get(user2)
	if has {
		fmt.Printf("user2:%v\n", user2)
	}

	// 使用 struct 结构体作为条件搜索
	user3 := &User{Id: 5}
	has, _ = engine.Get(user3)
	if has {
		fmt.Printf("user3:%v\n", user3)
	}

	user4 := &User{Name: "xxx"}
	has, _ = engine.Get(user4)
	if has {
		fmt.Printf("user4:%v\n", user4)
	}
}
