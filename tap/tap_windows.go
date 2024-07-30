package tap

import (
	"fmt"

	"github.com/songgao/water"
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
