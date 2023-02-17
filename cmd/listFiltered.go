/*
Copyright © 2023 Alexandre Moraes <alcmoraes89@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/alcmoraes/yip/routers/claro"
	"github.com/spf13/cobra"
)

// listFilteredCmd represents the listFiltered command
var listFilteredCmd = &cobra.Command{
	Use:   "listFiltered",
	Short: "List the devices that are currently blacklisted via mac address",
	Run: func(cmd *cobra.Command, args []string) {

		router := claro.NewClaroRouter()

		filteredDevices, err := router.GetFilteredDevices()
		if err != nil {
			panic(err)
		}

		fmt.Printf("===== DEVICES =====\n")
		for _, device := range filteredDevices {
			fmt.Printf("%s - %s\n", device.MacAddress, device.Name)
		}
		fmt.Printf("===================\n")

	},
}

func init() {
	rootCmd.AddCommand(listFilteredCmd)
}
