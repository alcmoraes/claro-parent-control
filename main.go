package main

import (
	"fmt"

	"github.com/alcmoraes/claro-parent-control/routers"
)

func main() {

	claro := routers.NewClaroRouter()

	claro.RetrieveToken()

	devices := []string{"64:A2:00:01:E8:6B", "CC:60:C8:31:26:E8"}

	// LOCK DEVICES
	for _, d := range devices {
		if ok := claro.FilterDeviceByMac(d); ok {
			fmt.Printf("Device %s locked!\n", d)
		}
	}

	// UNLOCK DEVICES
	// for _, d := range devices {
	// 	if ok := claro.UnfilterDeviceByMac(d); ok {
	// 		fmt.Printf("Device %s unlocked!\n", d)
	// 	}
	// }

}