package tap

import (
	"fmt"
	"log"
	"net"

	"github.com/songgao/water"
	"github.com/yudaiyan/gonetsh/netsh"
	"k8s.io/utils/exec"
)

func CreateTap() (*water.Interface, error) {
	config := water.Config{
		DeviceType: water.TAP,
		PlatformSpecificParams: water.PlatformSpecificParams{
			ComponentID: "tap0901",
		},
	}
	iface, err := water.New(config)
	if err != nil {
		return nil, fmt.Errorf("create tap %s", err.Error())
	}
	return iface, nil
}

func DeleteTap(name string) error {
	return noImplemention()
}

func AddrAdd(name string, addr string) error {
	h := netsh.New(exec.New())
	return h.SetIPAddress(name, addr)
}

func noImplemention() error {
	log.Println("no implemention")
	return nil
}

func LinkSetHardwareAddr(name string, mac net.HardwareAddr) error {
	return noImplemention()
}

func LinkSetUp(name string) error {
	return noImplemention()
}

func LinkDel(name string) error {
	return noImplemention()
}
