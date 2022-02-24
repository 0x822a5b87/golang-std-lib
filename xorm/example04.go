package main

import "fmt"

// Get()方法只能返回单条记录，其生成的 SQL 语句总是有LIMIT 1
func example04() {
	users := make([]User, 0)
	fmt.Println("len = ", len(users), ", capacity = ", cap(users))
	// len =  0 , capacity =  0
	err := engine.Where("age > ? and age < ?", 10, 100).Find(&users)
	if err != nil {
		return
	}
	for _, user := range users {
		fmt.Println("user.Name = ", user.Name, " user.Age = ", user.Age)
		// user.Name =  xxx  user.Age =  20
	}

	mapUsers := make(map[int64]User)
	// map的键为主键，所以如果表为复合主键就不能使用这种方式了。
	engine.Where("length(name) = ?", 3).Find(&mapUsers)
	for k, v := range mapUsers {
		fmt.Println("users whose has name of length 3: key = ", k, ", v = ", v)
		// users whose has name of length 3: key =  1 , v =  {1 xxx xxx 20 123456 2022-02-24 10:28:25 +0800 CST 2022-02-24 10:28:25 +0800 CST}
	}
}
