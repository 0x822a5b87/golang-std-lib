package main

import (
	"flag"
)

const (
	DefaultBid       = "b_ieg_idata_cdp_tlog3_test"
	DefaultTid       = "iDataLogMessageTest"
	DefaultSchema    = "__tablename,jobName,dtEventTime"
	DefaultBatch     = 10
	DefaultHost      = `{"cluster_id":95,"address":[{"port":"46801","host":"11.149.51.104","id":"1"},{"port":"46801","host":"11.149.51.109","id":"2"},{"port":"46801","host":"11.149.51.64","id":"3"},{"port":"46801","host":"11.149.51.91","id":"4"},{"port":"46801","host":"11.149.51.94","id":"5"},{"port":"46801","host":"30.44.132.123","id":"6"},{"port":"46801","host":"30.44.132.239","id":"7"},{"port":"46801","host":"30.44.132.70","id":"8"},{"port":"46801","host":"30.44.132.88","id":"9"}],"size":9,"load":25,"isInterVisit":1,"net_tag":"all","switch":0}`
	DefaultMessage   = `iDataLogMessageTest|tglog|`
	DefaultSleepTime = 2
)

func init() {
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	flag.StringVar(&bid, "bid", DefaultBid, "bid")
	flag.StringVar(&tid, "tid", DefaultTid, "tid")
	flag.StringVar(&schema, "schema", DefaultSchema, "schema")
	flag.IntVar(&batchSize, "batchSize", DefaultBatch, "batch size")
	flag.StringVar(&defaultHost, "defaultHost", DefaultHost, "default host")
	flag.StringVar(&message, "message", DefaultMessage, "default host")
	flag.DurationVar(&sleep, "sleep", DefaultSleepTime, "default sleep time")

	flag.Parse()
}
