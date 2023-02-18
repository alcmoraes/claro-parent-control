package claro

import (
	"github.com/alcmoraes/yip/routers"
	"github.com/go-resty/resty/v2"
)

type RefreshTokenResponse struct {
	AccessToken string `json:"accessToken"`
}

type FilteredDevicesResponse struct {
	MaxRules int `json:"maxRules"`
	Rules    []routers.FilteredDevice
}

type FilterRule struct {
	Active     bool   `json:"active"`
	MacAddress string `json:"macAddress"`
}

type ClaroRouter struct {
	Username   string
	Password   string
	Token      string
	HTTPClient *resty.Client
	Routes     map[string]string
}
