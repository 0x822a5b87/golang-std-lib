package context

import (
	"context"
	"fmt"
	"github.com/pborman/uuid"
	"strconv"
	"strings"
	"time"
)

const (
	KEY = "trace_id"
)

func NewRequestID() string {
	return strings.Replace(uuid.New(), "-", "", -1)
}

func NewContextWithTraceID() context.Context {
	ctx := context.WithValue(context.Background(), KEY, NewRequestID())
	for i := 0; i < 5; i++ {
		ctx = context.WithValue(ctx, strconv.Itoa(i), strconv.Itoa(i))
		fmt.Printf("key = %s, value = %s\n", strconv.Itoa(i), strconv.Itoa(i))
	}
	return ctx
}

func PrintLog(ctx context.Context, message string) {
	fmt.Printf("%s|info|trace_id=%s|%s\n", time.Now().Format("2006-01-02 15:04:05"), GetContextValue(ctx, KEY), message)
	for i := 0; i < 5; i++ {
		fmt.Printf("key = %s, value = %s\n", strconv.Itoa(i), GetContextValue(ctx, strconv.Itoa(i)))
	}
}

func GetContextValue(ctx context.Context, k string) string {
	v, ok := ctx.Value(k).(string)
	if !ok {
		return ""
	}
	return v
}

func ProcessEnter(ctx context.Context) {
	PrintLog(ctx, "Golang梦工厂")
}
