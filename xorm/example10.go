package main

import "fmt"

func example10() {
	user := User{}
	_, err := engine.Where("id = 1").Get(&user)
	if err != nil {
		return
	}
	logger.Info(fmt.Sprint("user = ", user))
	// {"level":"info","ts":1645692041.609569,"caller":"xorm/example10.go:13","msg":"user = {1 xxx xxx 20 123456 2022-02-24 15:58:15 +0800 CST 2022-02-24 16:38:30 +0800 CST}"}

	updateUser := User{
		Name: "xxx-update",
	}
	update, err := engine.ID(1).Update(updateUser)
	if err != nil {
		return
	}
	logger.Info(fmt.Sprint("update = ", update))

	user2 := User{}
	// 对于传入结构体指针的情况，xorm只会更新非空的字段。
	_, err = engine.Where("id = 1").Get(&user2)
	logger.Info(fmt.Sprint("user2 = ", user2))
	// {"level":"info","ts":1645692077.064068,"caller":"xorm/example10.go:26","msg":"user = {1 xxx-update xxx 20 123456 2022-02-24 15:58:15 +0800 CST 2022-02-24 16:40:41 +0800 CST}"}

	// 如果一定要更新空字段，需要使用Cols()方法显示指定更新的列。使用Cols()方法指定列后，即使字段为空也会更新：
	updateUser = User{
		Name: "xxx-update-with-cols",
	}
	engine.ID(1).Cols("Name", "salt").Update(updateUser)
	yes, err := engine.Where("id = 1").Get(&updateUser)
	logger.Info(fmt.Sprint("yes = ", yes, ", user = ", updateUser))
}
