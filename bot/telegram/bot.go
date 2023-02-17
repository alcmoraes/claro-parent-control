package telegram

import (
	"fmt"
	"log"
	"time"

	"github.com/alcmoraes/yip/routers/claro"
	"github.com/spf13/viper"
	tele "gopkg.in/telebot.v3"
)

type TelegramBot struct {
	bot    *tele.Bot
	router *claro.ClaroRouter
}

func (t *TelegramBot) Start() {
	pref := tele.Settings{
		Token:  viper.GetString("telegram.token"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)

	t.bot = b

	if err != nil {
		log.Fatal(err)
		return
	}

	// Authenticate (or re-authenticate) your session in the router
	b.Handle("/auth", t.AuthRouter)
	// Block a device by mac
	b.Handle("/block", t.FilterMac)
	// Unblocks a device by mac
	b.Handle("/unblock", t.UnfilterMac)

	b.Start()
}

func (t *TelegramBot) AuthRouter(c tele.Context) error {
	if err := t.router.RefreshToken(); err != nil {
		return c.Send("Failed to authenticate in the router")
	}
	return c.Send("Bot authenticated successfully")
}

func (t *TelegramBot) FilterMac(c tele.Context) error {
	devices, err := t.router.ListDevices()
	if err != nil {
		return c.Send("Failed to list devices")
	}

	menu := &tele.ReplyMarkup{}

	options := make([]tele.Row, 0)

	for _, d := range devices {
		btn := menu.Text(fmt.Sprintf("%s (%s)", d.MacAddress, d.Name))
		t.bot.Handle(&btn, func(c tele.Context) error {
			if err := t.router.FilterDeviceByMac(d.MacAddress); err != nil {
				return c.Send("Failed to block the device")
			} else {
				return c.Send("Device blocked successfully")
			}
		})
		options = append(options, menu.Row(btn))
	}
	menu.Reply(options...)

	return c.Send("Choose the device to block:", menu)
}

func (t *TelegramBot) UnfilterMac(c tele.Context) error {
	devices, err := t.router.GetFilteredDevices()
	if err != nil {
		return c.Send("Failed to list devices")
	}

	menu := &tele.ReplyMarkup{}

	options := make([]tele.Row, 0)

	for _, d := range devices {
		btn := menu.Text(d.MacAddress)
		t.bot.Handle(&btn, func(c tele.Context) error {
			if err := t.router.UnfilterDeviceByMac(d.MacAddress); err != nil {
				return c.Send("Failed to unblock the device")
			} else {
				return c.Send("Device unblocked successfully")
			}
		})
		options = append(options, menu.Row(btn))
	}
	menu.Reply(options...)

	return c.Send("Choose the device to unblock:", menu)
}

func NewTelegramBot() *TelegramBot {
	return &TelegramBot{
		router: claro.NewClaroRouter(),
	}
}
