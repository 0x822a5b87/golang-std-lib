package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func example01() {
	engine, err := xorm.NewEngine("mysql", "root:123456@/test?charset=utf8")
	engine.ShowSQL(true)

	if err != nil {
		logger.Fatal(fmt.Sprint(err))
	}

	// Sync2()方法会根据User的结构自动创建一个user表
	err = engine.Sync2(new(User))
	if err != nil {
		logger.Fatal(fmt.Sprint(err))
	}
}
