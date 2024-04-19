package main

import (
	"fmt"
	"net"
	"time"
)

func check(destination string, port string) string {
	address := destination + ":" + port
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable, \n Error: %v", address, err)
	} else {
		status = fmt.Sprintf("[UP] %v is reachable, \n from %v\n To: %v\n", destination, conn.LocalAddr(), conn.RemoteAddr())
	}

	return status
}
