package tap

import (
	"fmt"

	"github.com/songgao/water"
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
