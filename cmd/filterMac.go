/*
Copyright © 2023 Alexandre Moraes <alcmoraes89@gmail.com>
*/
package cmd

import (
	"github.com/alcmoraes/yip/routers/claro"
	"github.com/spf13/cobra"
)

// filterMacCmd represents the filterMac command
var filterMacCmd = &cobra.Command{
	Use:   "filterMac",
	Short: "Adds a device to the blacklist by a given mac address",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		router := claro.NewClaroRouter()

		for _, mac := range args {
			router.FilterDeviceByMac(mac)
		}
		
	},
}

func init() {
	rootCmd.AddCommand(filterMacCmd)
}
