/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/alcmoraes/yip/bot/telegram"
	"github.com/spf13/cobra"

	"github.com/alcmoraes/yip/routers/claro"
)

// telegramBotCmd represents the telegramBot command
var telegramBotCmd = &cobra.Command{
	Use:   "telegramBot",
	Short: "Starts a Telegram BOT to interact with the router",
	Run: func(cmd *cobra.Command, args []string) {
		bot := telegram.NewTelegramBot(claro.NewClaroRouter())
		bot.Start()
	},
}

func init() {
	rootCmd.AddCommand(telegramBotCmd)
}
