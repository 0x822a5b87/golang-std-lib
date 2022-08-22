package main

import (
	"fmt"
	"git.woa.com/pdata/common/tdbank"
	"git.woa.com/pdata/slog"
	"strings"
	"time"
)

var bid string
var tid string
var schema string
var batchSize int
var defaultHost string
var message string

var tube *tdbank.Reporter

func init() {
	tdbank.UseLog()
	tdbank.SetTableSchema(bid, tid, strings.Split(schema, ","))
	config := tdbank.Config{
		BusinessConfig: tdbank.BusinessConfig{
			BID:         bid,
			DefaultHost: defaultHost,
		},
		DialTimeout: time.Second * 10,
		IdleTimeout: time.Minute * 10,
		NetTimeout:  time.Second * 10,
		SendBatch:   batchSize,
	}
	var err error
	tube, err = tdbank.NewReporter(config)
	if err != nil {
		slog.Error(err, "初始化Reporter出错")
		panic("初始化Reporter出错")
	}
}

func ProcessSendRawMsg() {
	go processError()
	for i := 1; i <= 2*batchSize; i++ {
		tube.SendRawMsg(tid, []byte(message))
		fmt.Println("send message bid = [" + bid + "], tid = [" + tid + "], message = [" + message + "]")
	}
}

func processError() {
	messages := tube.Error()
	for m := range messages {
		fmt.Println("error message, tid = " + m.TID + ", Content = " + string(m.Content))
	}
}

func main() {
	fmt.Println("start tube sender!")
	ProcessSendRawMsg()
	fmt.Println("end tube sender!")
}
