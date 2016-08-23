package main

import (
	"flag"
	"github.com/golang/glog"
	"log"
)

func main() {
	flag.Parse()
	glog.CopyStandardLogTo("info")
	defer glog.Flush()

	log.Println("log.println")

	glog.Infof("info")
	glog.Warningf("warn")
	glog.Errorf("error")

	glog.V(0).Infof("v=0")
	glog.V(1).Infof("v=1")
	glog.V(2).Infof("v=2")
	glog.V(3).Infof("v=3")

	if glog.V(2) {
		glog.Infof("if v=2")
	}

	//	glog.Fatalf("Initialization failed: %s", fmt.Errorf("foo"))

}
