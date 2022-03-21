package testify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Mock 简单来说就是构造一个仿对象，仿对象提供和原对象一样的接口，在测试中用仿对象来替换原对象。
//这样我们可以在原对象很难构造，特别是涉及外部资源（数据库，访问网络等）时进行测试。
//例如，我们现在要编写一个从一个站点拉取用户列表信息的程序，拉取完成之后程序显示和分析。
//如果每次都去访问网络会带来极大的不确定性，甚至每次返回不同的列表，这就给测试带来了极大的困难。我们可以使用 Mock 技术。

type User struct {
	Name string
	Age  int
}

type ICrawler interface {
	GetUserList() ([]*User, error)
}

type MyCrawler struct {
	url string
}

func (c *MyCrawler) GetUserList() ([]*User, error) {
	resp, err := http.Get(c.url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var userList []*User
	err = json.Unmarshal(data, &userList)
	if err != nil {
		return nil, err
	}

	return userList, nil
}

func GetAndPrintUsers(crawler ICrawler) {
	users, err := crawler.GetUserList()
	if err != nil {
		return
	}

	for _, u := range users {
		_ = u.Age
	}
}
