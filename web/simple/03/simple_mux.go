package simple

import (
	"fmt"
	"net/http"
)

// 实现自定义的 Mux

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHelloRoute(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayHelloRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
