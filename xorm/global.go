package main

import (
	"fmt"
	"go.uber.org/zap"
	"math/rand"
	"time"
	"xorm.io/xorm"
)

var logger, _ = zap.NewProduction()

var engine, _ = xorm.NewEngine("mysql", "root:123456@/test?charset=utf8")

type User struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

type Sum struct {
	Id    int64
	Money int32
	Rate  float32
}

type Player struct {
	Id        int64
	Name      string
	Age       int
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	// 如果我们为time.Time类型的字段设置了xorm:"deleted"标签，删除数据时，只是设置删除时间，并不真正删除记录。
	DeletedAt time.Time `xorm:"deleted"`
}

func initUser() {
	users := []User{
		{
			Id:      1,
			Name:    "xxx",
			Salt:    "xxx",
			Age:     20,
			Passwd:  "123456",
			Created: time.Time{},
			Updated: time.Time{},
		},
		{
			Id:      2,
			Name:    "yyy",
			Salt:    "yyy",
			Age:     30,
			Passwd:  "yyy",
			Created: time.Time{},
			Updated: time.Time{},
		},
		{
			Id:      3,
			Name:    "zzz",
			Salt:    "zzz",
			Age:     40,
			Passwd:  "zzz",
			Created: time.Time{},
			Updated: time.Time{},
		},
	}

	u := &User{}
	for _, user := range users {
		has, _ := engine.ID(user.Id).Get(u)
		if !has {
			_, err := engine.InsertOne(user)
			if err != nil {
				return
			}
			logger.Info(fmt.Sprint("user : ", user))
		}
	}
}

func initSum() {
	engine.Sync2(&Sum{})

	var slice []*Sum
	for i := 0; i < 100; i++ {
		slice = append(slice, &Sum{
			Money: rand.Int31n(10000),
			Rate:  rand.Float32(),
		})
	}
	engine.Insert(&slice)
}

func init() {
	initUser()
	initSum()
}
