package simple

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

func StartHelloName() {
	http.HandleFunc("/", sayHelloName)       // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
