package messages

import "github.com/spf13/viper"

var i18n map[string]string

func init() {
	lang := viper.GetString("language")
	switch lang {
		case "br": i18n = Brazilian()
	default: i18n = English()
	}
}

func Get(s string) string {
	return i18n[s]
}