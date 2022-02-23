package main

import (
	"fmt"
	"go.uber.org/zap"
	"net"
)

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

// {"level":"info","msg":"hello world","ip":"192.168.255.10"}
func example07() {
	ip := GetOutboundIP()
	logger := zap.NewExample(zap.Fields(
		zap.String("ip", ip.String()),
	))

	logger.Info("hello world")
}
