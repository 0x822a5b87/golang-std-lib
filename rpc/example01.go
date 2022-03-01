package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

// Arith 定义 Arith
type Arith int

// Multiply 为 Arith 绑定 Multiply 方法
func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// Divide 为 Arith 绑定 Divide 方法
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by 0")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func example01() {
	go example01Server()
	time.Sleep(1 * time.Second)
	example01Client()
}

func example01Server() {
	arith := new(Arith)
	// Register publishes in the server the set of methods of the
	// receiver value that satisfy the following conditions:
	//	- exported method of exported type
	//	- two arguments, both of exported type
	//	- the second argument is a pointer
	//	- one return value, of type error
	// It returns an error if the receiver is not an exported type or has
	// no suitable methods. It also logs the error using package log.
	// The client accesses each method using a string of the form "Type.Method",
	// where Type is the receiver's concrete type.
	err := rpc.Register(arith)
	if err != nil {
		return
	}
	// 注册http路由
	rpc.HandleHTTP()
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("serve error:", err)
	}
}

func example01Client() {
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Multiply error:", err)
	}
	fmt.Printf("Multiply: %d*%d=%d\n", args.A, args.B, reply)

	args = &Args{15, 6}
	var quo Quotient
	err = client.Call("Arith.Divide", args, &quo)
	if err != nil {
		log.Fatal("Divide error:", err)
	}
	fmt.Printf("Divide: %d/%d=%d...%d\n", args.A, args.B, quo.Quo, quo.Rem)
}
