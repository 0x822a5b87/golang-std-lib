package main

import (
	"log"
	"net/http"
	"net/rpc"
)

func example03() {
	arith := new(Arith)
	server := rpc.NewServer()
	server.RegisterName("math", arith)
	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("serve error:", err)
	}
}
