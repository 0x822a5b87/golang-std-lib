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

var tdBankReport *tdbank.Reporter

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
	tdBankReport, err = tdbank.NewReporter(config)
	if err != nil {
		slog.Error(err, "初始化Reporter出错")
		panic("初始化Reporter出错")
	}
}

func ProcessSendRawMsg() {
	for i := 1; i <= batchSize; i++ {
		tdBankReport.SendRawMsg(tid, []byte(message))
		fmt.Println("send message, tid = " + tid + ", message = " + message)
	}
}

func main() {
	fmt.Println("start tube sender!")
	ProcessSendRawMsg()
	fmt.Println("end tube sender!")
}
