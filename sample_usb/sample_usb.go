package main

import (
	"github.com/google/gousb"
	"github.com/golang/glog"
	"flag"
)

func main() {
	flag.Parse()
	ctx := gousb.NewContext()
	defer ctx.Close()
	vid, pid := gousb.ID(1452), gousb.ID(34304)
	devs, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		glog.Infof("vendor %d product %d", desc.Vendor, desc.Product)
		return desc.Vendor == vid && desc.Product == pid
	})
	defer func() {
		for _, d := range devs {
			d.Close()
		}
	}()
	if err != nil {
		glog.Exitf("get devices failed: %v", err)
		return
	}
	glog.Infof("found %d devs", len(devs))
	for _, dev := range devs {
		{
			text := dev.String()
			glog.Infof("String: %s", text)
		}
		{
			text, err := dev.Manufacturer()
			glog.Infof("Manufacturer: %s %v", text, err)
		}
		{
			text, err := dev.Product()
			glog.Infof("Product: %s %v", text, err)
		}
		{
			text, err := dev.SerialNumber()
			glog.Infof("SerialNumber: %s %v", text, err)
		}
		intf, done, err := dev.DefaultInterface()
		if err != nil {
			glog.Exitf("%s.DefaultInterface(): %v", dev, err)
		}
		defer done()
		epIn, err := intf.InEndpoint(6)
		if err != nil {
			glog.Exitf("%s.InEndpoint(6): %v", intf, err)
		}
		buf := make([]byte, 10*epIn.Desc.MaxPacketSize)
		for {
			n, err := epIn.Read(buf)
			if err != nil {
				glog.Exitf("read failed: %v", err)
			}
			glog.Infof("read %d", n)
		}
	}
}
