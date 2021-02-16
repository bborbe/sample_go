package main

import (
	"fmt"
	"github.com/foomo/htpasswd"
	"github.com/golang/glog"
)

func main() {
	file := "/tmp/htpasswd"
	name := "foo"
	password := "bar"
	err := htpasswd.SetPassword(file, name, password, htpasswd.HashBCrypt)
	if err != nil {
		glog.Exit(err)
	}
	fmt.Printf("done\n")
}
