package main

import "fmt"

func example09() {
	user := &User{Name: "lzy", Age: 50}

	affected, _ := engine.Insert(user)
	// 单条数据写入时，user.Id 会被自动赋值
	logger.Info(fmt.Sprintf("%d records inserted, user.id:%d\n", affected, user.Id))

	users := make([]*User, 2)
	users[0] = &User{Name: "xhq", Age: 41}
	users[1] = &User{Name: "lhy", Age: 12}

	affected, _ = engine.Insert(&users)
	// 批量插入时，自增长主键 Id 不会被自动赋值，所以 users[0].Id 和 users[1].Id 都是 0
	logger.Info(fmt.Sprintf("%d records inserted, id1:%d, id2:%d", affected, users[0].Id, users[1].Id))
	// {"level":"info","ts":1645691462.182535,"caller":"xorm/example09.go:16","msg":"2 records inserted, id1:0, id2:0"}
}
