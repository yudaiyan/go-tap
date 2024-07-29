package tap

import (
	"fmt"
	"log"
	"net"

	"github.com/songgao/water"
	"github.com/vishvananda/netlink"
	"golang.org/x/exp/rand"
)

func RandStr(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func CreateTap() (*water.Interface, error) {
	config := water.Config{
		DeviceType: water.TAP,
	}
	config.Name = fmt.Sprintf("tap-%s", RandStr(4))

	iface, err := water.New(config)
	if err != nil {
		return nil, fmt.Errorf("create tap %s", err.Error())
	}
	return iface, nil
}

func DeleteTap(name string) error {
	link, _ := netlink.LinkByName(name)
	if link != nil {
		return netlink.LinkDel(link)
	}
	log.Printf("tap %s not found", name)
	return nil
}

func AddrAdd(name string, addr string) error {
	addr2, err := netlink.ParseAddr(addr)
	if err != nil {
		return err
	}
	link, err := netlink.LinkByName(name)
	if err != nil {
		return err
	}
	return netlink.AddrAdd(link, addr2)
}

func LinkSetHardwareAddr(name string, mac net.HardwareAddr) error {
	link, _ := netlink.LinkByName(name)
	if err := netlink.LinkSetHardwareAddr(link, mac); err != nil {
		return err
	}
	return nil
}

func LinkSetUp(name string) error {
	link, _ := netlink.LinkByName(name)
	return netlink.LinkSetUp(link)
}

func LinkDel(name string) error {
	link, _ := netlink.LinkByName(name)
	return netlink.LinkDel(link)
}
