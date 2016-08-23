package main

import (
	"os"

	"github.com/coreos/pkg/capnslog"
)

var plog = capnslog.NewPackageLogger("github.com/bborbe/sample_go", "sample_plog")

func main() {

	capnslog.SetGlobalLogLevel(capnslog.DEBUG)
	capnslog.SetFormatter(capnslog.NewGlogFormatter(os.Stderr))
	f, err := capnslog.NewJournaldFormatter()
	if err == nil {
		capnslog.SetFormatter(f)
	}

	plog.Tracef("trace")
	plog.Debugf("debug")
	plog.Noticef("notice")
	plog.Infof("info")
	plog.Warningf("warn")
	plog.Errorf("error")
	plog.Fatalf("fatal")
	plog.Panicf("panic")
}
