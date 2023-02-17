package routers

import (
	"encoding/json"
	"os"

	"github.com/go-resty/resty/v2"
)

type RetrieveTokenResponse struct {
	AccessToken string `json:"accessToken`
}

type FilteredDevicesResponse struct {
	MaxRules int `json:"maxRules"`
	Rules    []FilteredDevice
}

type FilteredDevice struct {
	Id         int    `json:"id"`         // : 0,
	Active     bool   `json:"active"`     // : true,
	Name       string `json:"name"`       // : "cpe-bridgefilter-37",
	MacAddress string `json:"macAddress"` // : "64:A2:00:01:E8:6B",
	DeviceId   int    `json:"deviceId"`   // : 0
}

type FilterRule struct {
	Active     bool   `json:"active"`
	MacAddress string `json:"macAddress"`
}

type Device struct {
	Id                   string `json:"id"`                   // : "64:64:4A:39:E3:38",
	Connected            string `json:"connected"`            // : true,
	Type                 string `json:"type"`                 // : "PC",
	Interface            string `json:"interface"`            // : "lan",
	Path                 string `json:"path"`                 // : "2",
	MacAddress           string `json:"macAddress"`           // : "64:64:4A:39:E3:38",
	IpAddress            string `json:"ipAddress"`            // : "192.168.0.43",
	Name                 string `json:"name"`                 // : "MiWiFi-R4CM",
	LeaseTime            string `json:"leaseTime"`            // : 3600,
	ExpireTime           string `json:"expireTime"`           // : "--- --- -- --:--:-- ----",
	Ipv6Address          string `json:"ipv6Address"`          // : "",
	Ipv6LinkLocalAddress string `json:"ipv6LinkLocalAddress"` // : "",
	LastConnected        string `json:"lastConnected"`        // : "",
	Rssi                 string `json:"rssi"`                 // : 0,
	OperatingStandard    string `json:"operatingStandard"`    // : "",
	Age                  string `json:"age"`                  // : 0
}

type ClaroRouter struct {
	Username   string
	Password   string
	Token      string
	HTTPClient *resty.Client
	Routes     map[string]string
}

func (c *ClaroRouter) RetrieveToken() {
	resp, err := c.HTTPClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{
			"userName": c.Username,
			"password": c.Password,
		}).
		Post(c.Routes["login"])
	if err != nil {
		panic(err)
	}

	var response RetrieveTokenResponse
	json.Unmarshal(resp.Body(), &response)
	c.Token = response.AccessToken
}

func (c *ClaroRouter) ListDevices() []Device {
	resp, err := c.HTTPClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Access-Token", c.Token).
		Get(c.Routes["devices"])
	if err != nil {
		panic(err)
	}

	var response []Device
	json.Unmarshal(resp.Body(), &response)
	return response
}

func (c *ClaroRouter) GetFilteredDevices() (output []FilteredDevice) {
	resp, err := c.HTTPClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Access-Token", c.Token).
		Get(c.Routes["macFiltering"])
	if err != nil {
		panic(err)
	}

	var response FilteredDevicesResponse
	json.Unmarshal(resp.Body(), &response)
	return response.Rules
}

func (c *ClaroRouter) FilterDeviceByMac(mac string) bool {
	lockedDevices := c.GetFilteredDevices()
	for _, device := range lockedDevices {
		if device.MacAddress == mac {
			return true
		}
	}

	resp, err := c.HTTPClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Access-Token", c.Token).
		SetBody(map[string]interface{}{
			"macAddress": mac,
			"active":     true,
		}).
		Post(c.Routes["macFiltering"])
	if err != nil {
		panic(err)
	}

	return resp.IsSuccess()
}

func (c *ClaroRouter) UnfilterDeviceByMac(mac string) bool {
	newRules := make([]FilterRule, 0)
	lockedDevices := c.GetFilteredDevices()
	for _, device := range lockedDevices {
		if device.MacAddress != mac {
			newRules = append(newRules, FilterRule{
				Active:     true,
				MacAddress: device.MacAddress,
			})
		}
	}

	resp, err := c.HTTPClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Access-Token", c.Token).
		SetBody(map[string]interface{}{
			"rules": newRules,
		}).
		Put(c.Routes["macFiltering"])
	if err != nil {
		panic(err)
	}

	return resp.IsSuccess()
}

func NewClaroRouter() *ClaroRouter {
	return &ClaroRouter{
		HTTPClient: resty.New().SetHostURL(
			os.Getenv("CLARO_ROUTER_HOST"),
		),
		Username:   os.Getenv("CLARO_ROUTER_USERNAME"),
		Password:   os.Getenv("CLARO_ROUTER_PASSWORD"),
		Routes: map[string]string{
			"login":        "api/v1/gateway/users/login",
			"devices":      "api/v1/gateway/devices?connected=true",
			"macFiltering": "api/v1/service/macFiltering",
		},
	}
}
