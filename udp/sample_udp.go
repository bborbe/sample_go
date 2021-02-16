package main

import (
	"github.com/golang/glog"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:1337")
	if err != nil {
		glog.Exit(err)
	}
	conn.Write([]byte("hello"))
}
