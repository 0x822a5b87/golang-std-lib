package suite

import (
	"fmt"
	"github.com/stretchr/testify/suite"
)

//testify提供了测试套件的功能（TestSuite），testify测试套件只是一个结构体，内嵌一个匿名的suite.Suite结构。
//测试套件中可以包含多个测试，它们可以共享状态，还可以定义钩子方法执行初始化和清理操作。
//钩子都是通过接口来定义的，实现了这些接口的测试套件结构在运行到指定节点时会调用对应的方法。
//有以下常用钩子函数
//SetupSuite		在所有测试开始之前运行
//TearDownSuite		在所有测试结束之后运行
//SetupTest			在每个测试开始之前运行
//TearDownTest		在每个测试结束之后运行
//BeforeTest/AfterTest，它们分别在每个测试运行前/后调用，接受套件名和测试名作为参数

type MyTestSuit struct {
	suite.Suite
	testCount uint32
}

func (s *MyTestSuit) SetupSuite() {
	fmt.Println("SetupSuite")
	go startServer()
}

func (s *MyTestSuit) TearDownSuite() {
	fmt.Println("TearDownSuite")
}

func (s *MyTestSuit) SetupTest() {
	fmt.Printf("SetupTest test count:%d\n", s.testCount)
}

func (s *MyTestSuit) TearDownTest() {
	s.testCount++
	fmt.Printf("TearDownTest test count:%d\n", s.testCount)
}

func (s *MyTestSuit) BeforeTest(suiteName, testName string) {
	fmt.Printf("BeforeTest suite:%s test:%s\n", suiteName, testName)
}

func (s *MyTestSuit) AfterTest(suiteName, testName string) {
	fmt.Printf("AfterTest suite:%s test:%s\n", suiteName, testName)
}

func (s *MyTestSuit) TestExample() {
	fmt.Println("TestExample")
	if server != nil {
		server.Close()
	}
}
