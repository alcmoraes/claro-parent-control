package telegram

import (
	"fmt"
	"log"
	"strings"
	"time"

	"golang.org/x/sys/unix"

	"github.com/alcmoraes/yip/messages"
	"github.com/alcmoraes/yip/routers"
	"github.com/spf13/viper"
	tele "gopkg.in/telebot.v3"
)

type TelegramBot struct {
	bot       *tele.Bot
	router    routers.Router
	authGroup map[int64]bool
	nameCache map[string]string
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

	menu := &tele.ReplyMarkup{ForceReply: true, ResizeKeyboard: true}

	btnReauthenticateRouter := menu.Text(messages.Get("BTN_REAUTH_ROUTER"))
	btnWhitelistMac := menu.Text(messages.Get("BTN_ALLOW_MAC"))
	btnBlacklistMac := menu.Text(messages.Get("BTN_BLOCK_MAC"))
	btnClearBlacklist := menu.Text(messages.Get("BTN_CLEAR_BLACKLIST"))
	btnRestart := menu.Text(messages.Get("BTN_RESTART"))

	menu.Reply(
		menu.Row(btnReauthenticateRouter),
		menu.Row(btnBlacklistMac, btnWhitelistMac),
		menu.Row(btnClearBlacklist, btnRestart),
	)

	// Login to bot
	b.Handle("/start", func(c tele.Context) error {
		return c.Send(messages.Get("MSG_WELCOME"), menu)
	})

	b.Handle("/login", t.Login)
	b.Handle(&btnReauthenticateRouter, t.ReauthenticateRouter, t.LockMiddleware)
	b.Handle(&btnWhitelistMac, t.WhitelistMac, t.LockMiddleware)
	b.Handle(&btnBlacklistMac, t.BlacklistMac, t.LockMiddleware)
	b.Handle(&btnClearBlacklist, t.ClearBlacklist, t.LockMiddleware)
	b.Handle(&btnRestart, t.RestartPi, t.LockMiddleware)

	b.Handle(tele.OnCallback, t.OnCallback, t.LockMiddleware)

	b.Start()
}

func (t *TelegramBot) RestartPi(c tele.Context) error {
	err := unix.Reboot(unix.LINUX_REBOOT_CMD_RESTART);
	return err
}

func (t *TelegramBot) LockMiddleware(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		if c.Callback() != nil {
			defer c.Respond()
		}
		if t.authGroup[c.Sender().ID] {
			return next(c) // continue execution chain
		}
		return c.Send(messages.Get("MSG_UNAUTHENTICATED"))
	}
}

func (t *TelegramBot) Login(c tele.Context) error {
	arguments := c.Args()
	if len(arguments) == 0 {
		return c.Send(messages.Get("MSG_UNAUTHENTICATED"))
	}
	if arguments[0] == viper.GetString("telegram.password") {
		t.authGroup[c.Sender().ID] = true
		return c.Send(messages.Get("MSG_AUTHENTICATED"))
	} else {
		return c.Send(messages.Get("MSG_WRONG_PASSWORD"))
	}
}

func (t *TelegramBot) ReauthenticateRouter(c tele.Context) error {
	if err := t.router.RefreshToken(); err != nil {
		return c.Send(messages.Get("MSG_REAUTH_FAILED"))
	}
	return c.Send(messages.Get("MSG_REAUTH_SUCCESS"))
}

func (t *TelegramBot) BlacklistMac(c tele.Context) error {
	devices, err := t.router.ListDevices()
	if err != nil {
		return c.Send(messages.Get("MSG_FAILED_TO_GET_BLACKLIST"))
	}

	menu := &tele.ReplyMarkup{}
	options := make([]tele.Row, 0)

	for _, d := range devices {
		t.nameCache[strings.ToUpper(d.MacAddress)] = d.Name
		btn := menu.Data(fmt.Sprintf("%s (%s)", d.MacAddress, d.Name), "/block", d.MacAddress)
		options = append(options, menu.Row(btn))
	}
	menu.Inline(options...)

	return c.Send(messages.Get("MSG_CHOOSE_DEVICE"), menu)
}

func (t *TelegramBot) WhitelistMac(c tele.Context) error {
	devices, err := t.router.GetFilteredDevices()
	if err != nil {
		return c.Send(messages.Get("MSG_FAILED_TO_GET_BLACKLIST"))
	}
	if len(devices) == 0 {
		return c.Send(messages.Get("MSG_NO_DEVICES_BLOCKED"))
	}

	menu := &tele.ReplyMarkup{}
	options := make([]tele.Row, 0)

	for _, d := range devices {
		cachedName := t.nameCache[strings.ToUpper(d.MacAddress)]
		if cachedName != "" {
			d.Name = cachedName
		}
		btn := menu.Data(fmt.Sprintf("%s (%s)", d.MacAddress, d.Name), "/unblock", d.MacAddress)
		options = append(options, menu.Row(btn))
	}
	menu.Inline(options...)

	return c.Send(messages.Get("MSG_CHOOSE_DEVICE"), menu)
}

func (t *TelegramBot) ClearBlacklist(c tele.Context) error {
	if err := t.router.ClearMacFilters(); err != nil {
		return c.Send(messages.Get("MSG_FAILED_TO_CLEAR_BLACKLIST"))
	}
	return c.Send(messages.Get("MSG_BLACKLIST_CLEARED"))
}

func (t *TelegramBot) OnCallback(c tele.Context) error {

	command := strings.TrimSpace(c.Args()[0])

	switch command {
	case "/unblock":
		if err := t.router.UnfilterDeviceByMac(c.Args()[1]); err != nil {
			return c.Send(messages.Get("MSG_FAILED_TO_UNBLOCK_DEVICE"))
		} else {
			return c.Send(messages.Get("MSG_DEVICE_UNBLOCKED"))
		}
	case "/block":
		if err := t.router.FilterDeviceByMac(c.Args()[1]); err != nil {
			return c.Send(messages.Get("MSG_FAILED_TO_BLOCK_DEVICE"))
		} else {
			return c.Send(messages.Get("MSG_DEVICE_BLOCKED"))
		}
	}

	return c.Send(messages.Get("MSG_UNKNOWN_COMMAND"))
}

func NewTelegramBot(r routers.Router) *TelegramBot {
	return &TelegramBot{
		router:    r,
		authGroup: make(map[int64]bool, 0),
		nameCache: make(map[string]string, 0),
	}
}
