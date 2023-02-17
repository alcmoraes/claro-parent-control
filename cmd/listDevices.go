/*
Copyright © 2023 Alexandre Moraes <alcmoraes89@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/alcmoraes/yip/routers/claro"
	"github.com/spf13/cobra"
)

// listDevicesCmd represents the listDevices command
var listDevicesCmd = &cobra.Command{
	Use:   "listDevices",
	Short: "List devices connected to the DHCP server",
	Run: func(cmd *cobra.Command, args []string) {
		
		router := claro.NewClaroRouter()

		devices := router.ListDevices()

		fmt.Printf("===== DEVICES =====\n")
		for _, device := range devices {
			fmt.Printf("%s - %s\n", device.MacAddress, device.Name)
		}
		fmt.Printf("===================\n")
	},
}

func init() {
	rootCmd.AddCommand(listDevicesCmd)
}
