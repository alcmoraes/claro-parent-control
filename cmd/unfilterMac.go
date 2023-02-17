/*
Copyright © 2023 Alexandre Moraes <alcmoraes89@gmail.com>
*/
package cmd

import (
	"github.com/alcmoraes/yip/routers/claro"
	"github.com/spf13/cobra"
)

// unfilterMacCmd represents the unfilterMac command
var unfilterMacCmd = &cobra.Command{
	Use:   "unfilterMac",
	Short: "Removes a device from the blacklist by a given mac address",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		router := claro.NewClaroRouter()

		for _, mac := range args {
			router.UnfilterDeviceByMac(mac)
		}
	},
}

func init() {
	rootCmd.AddCommand(unfilterMacCmd)
}
