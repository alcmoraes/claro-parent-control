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
	authGroup map[int64]bool
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

	// Login to bot
	b.Handle("/login", t.Login)
	// Authenticate (or re-authenticate) your session in the router
	b.Handle("/reauth", t.AuthRouter, t.LockMiddleware)
	// Block a device by mac
	b.Handle("/block", t.FilterMac, t.LockMiddleware)
	// Unblocks a device by mac
	b.Handle("/unblock", t.UnfilterMac, t.LockMiddleware)

	b.Start()
}

func (t *TelegramBot) Login(c tele.Context) error {
	arguments := c.Args()
	if len(arguments) == 0 {
		return c.Send("Please use /login [PASSWORD] to authenticate")
	}
	if arguments[0] == viper.GetString("telegram.password") {
		t.authGroup[c.Sender().ID] = true
		return c.Send("Login successful")
	} else {
		return c.Send("Wrong password")
	}
}

func (t *TelegramBot) LockMiddleware(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		if c.Callback() != nil {
			defer c.Respond()
		}
		if t.authGroup[c.Sender().ID] {
			return next(c) // continue execution chain
		}
		return c.Send("You are not authenticated. Please use /login [PASSWORD] to authenticate")
	}
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
		
		func (b *tele.Btn, m string) {
			t.bot.Handle(b, func(c tele.Context) error {
				fmt.Println("Blocking device: ", m)
				if err := t.router.FilterDeviceByMac(m); err != nil {
					return c.Send("Failed to block the device")
				} else {
					return c.Send("Device blocked successfully")
				}
			})
		}(&btn, d.MacAddress)

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
		func (b *tele.Btn, m string) {
			t.bot.Handle(b, func(c tele.Context) error {
				fmt.Println("Unblocking device: ", m)
				if err := t.router.UnfilterDeviceByMac(m); err != nil {
					return c.Send("Failed to unblock the device")
				} else {
					return c.Send("Device unblocked successfully")
				}
			})
		}(&btn, d.MacAddress)
		options = append(options, menu.Row(btn))
	}
	menu.Reply(options...)

	return c.Send("Choose the device to unblock:", menu)
}

func NewTelegramBot() *TelegramBot {
	return &TelegramBot{
		router: claro.NewClaroRouter(),
		authGroup: make(map[int64]bool, 0),
	}
}
