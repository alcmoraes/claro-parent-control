package routers

import (
	"strings"

	"github.com/spf13/viper"
)

func RemoveWhitelisted(devices []Device) (output []Device) {

	if len(viper.GetStringSlice("whitelist")) == 0 {
		return devices
	}

	for _, device := range devices {
		isWhitelisted := false
		for _, mac := range viper.GetStringSlice("whitelist") {
			if strings.EqualFold(device.MacAddress, mac) {
				isWhitelisted = true
			}
		}
		if !isWhitelisted {
			output = append(output, device)
		}

	}
	return output
}

func IsWhitelisted(mac string) bool {
	for _, device := range viper.GetStringSlice("whitelist") {
		if strings.EqualFold(device, mac) {
			return true
		}
	}
	return false
}
