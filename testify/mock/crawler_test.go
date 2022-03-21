package testify

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

//使用 mock 的方法分为以下几个步骤
// 1. 实现我们需要mock的方法并绑定到 mock.Mock
// 		1.1 在方法的实现中，我们必须调用 Called() 方法，这个方法告诉 mock 一个 method 被调用了
//		1.2 其次我们必须调用 Get()/Int()/String()/Bool()/Error() 方法来定义一个返回值，并且 Get()/... 的参数就是返回值的顺序
// 2. 调用 mock.On(methodName).Return() 方法指定在 methodName 调用时的返回值
// 3.

type MockCrawler struct {
	mock.Mock
}

func (m *MockCrawler) GetUserList() ([]*User, error) {
	//实现GetUserList()方法时，需要调用Mock.Called()方法，传入参数（示例中无参数）。
	//Called()会返回一个mock.Arguments对象，该对象中保存着返回的值
	//它提供了对基本类型和error的获取方法Int()/String()/Bool()/Error()，
	//和通用的获取方法Get()，通用方法返回interface{}，需要类型断言为具体类型，它们都接受一个表示索引的参数。
	args := m.Called()
	return args.Get(0).([]*User), args.Error(1)
}

var (
	MockUsers []*User
)

func init() {
	MockUsers = append(MockUsers, &User{"xxx", 18})
	MockUsers = append(MockUsers, &User{"yyy", 20})
}

func TestGetUserList(t *testing.T) {
	crawler := new(MockCrawler)

	//crawler.On("GetUserList").Return(MockUsers, nil)是 Mock 发挥魔法的地方，
	//这里指示调用GetUserList()方法的返回值分别为MockUsers和nil，
	//返回值在上面的GetUserList()方法中被Arguments.Get(0)和Arguments.Error(1)获取。
	crawler.On("GetUserList").Return(MockUsers, nil)

	GetAndPrintUsers(crawler)

}
