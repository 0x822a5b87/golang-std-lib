# go web 开发

## reference

[Go web 编程](https://learnku.com/docs/build-web-application-with-golang)

## Go Web 编程

> 先看一个最简单的web服务器代码，步骤分为：
>
> 1. 声明 hadler function `sayHelloName`，这个函数用于处理输入并产生输出；
> 2. 将声明的函数 `sayHelloName` 绑定到路由 `/`
> 3. 启动web服务器并监听端口 `9090`

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello astaxie!")
	if err != nil {
		return 
	}
}

func startHelloName() {
	http.HandleFunc("/", sayHelloName)       // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
```

### 03.0 web 基础

#### web server 分析

> 在上面的web服务器中，有几个关键的组件：
>
> 1. Handler ：handler 负责回应一个 http 请求；
> 2. Server : A Server defines parameters for running an HTTP server.
>
> `Server` 包含了 Handler，

```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

```go
// A Server defines parameters for running an HTTP server.
// The zero value for Server is a valid configuration.
type Server struct {
    // ...
    Handler Handler // handler to invoke, http.DefaultServeMux if nil
    //...
}
```

```go
// ListenAndServe listens on the TCP network address addr and then calls
// Serve with handler to handle requests on incoming connections.
// Accepted connections are configured to enable TCP keep-alives.
//
// The handler is typically nil, in which case the DefaultServeMux is used.
//
// ListenAndServe always returns a non-nil error.
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
```

#### ServeMux 的自定义

> 我们前面小节讲述 conn.server 的时候，其实内部是调用了 http 包默认的路由器，通过路由器把本次请求的信息传递到了后端的处理函数。那么这个路由器是怎么实现的呢？

```go
// ServeMux is an DNS request multiplexer. It matches the zone name of
// each incoming request against a list of registered patterns add calls
// the handler for the pattern that most closely matches the zone name.
//
// ServeMux is DNSSEC aware, meaning that queries for the DS record are
// redirected to the parent zone (if that is also registered), otherwise
// the child gets the query.
//
// ServeMux is also safe for concurrent access from multiple goroutines.
//
// The zero ServeMux is empty and ready for use.
type ServeMux struct {
	z map[string]Handler
	m sync.RWMutex
}
```

```go
// Handler is implemented by any value that implements ServeDNS.
type Handler interface {
	ServeDNS(w ResponseWriter, r *Msg)
}
```

### 04.0 表单

> 表单是一个包含表单元素的区域。表单元素是允许用户在表单中（比如：文本域、下拉列表、单选框、复选框等等）输入信息的元素。表单使用表单标签（\<form>）定义。

```html
<form>
...
input 元素
...
</form>
```

#### 04.1. 处理表单的输入

> login.gtpl

```html
<html>
<head>
<title></title>
</head>
<body>
<form action="/login" method="post">
    用户名:<input type="text" name="username">
    密码:<input type="password" name="password">
    <input type="submit" value="登录">
</form>
</body>
</html>
```

```go
package simple

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // 获取请求的方法
	if r.Method == "GET" {
		t, err := template.ParseFiles("simple/04/login.gtpl")
		if err != nil {
			log.Println("err: ", err)
			return
		} else {
			log.Println(t.Execute(w, nil))
		}
	} else {
		err := r.ParseForm()   // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
		if err != nil {
			// handle error http.Error() for example
			log.Fatal("ParseForm: ", err)
		}
		// 请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func Start() {
	http.HandleFunc("/login", login)         // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
```



































