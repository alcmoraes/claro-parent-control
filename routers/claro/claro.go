package claro

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

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
	router := &ClaroRouter{
		HTTPClient: resty.New().SetHostURL(viper.GetString("claro.host")),
		Username:   viper.GetString("claro.username"),
		Password:   viper.GetString("claro.password"),
		Routes: map[string]string{
			"login":        "api/v1/gateway/users/login",
			"devices":      "api/v1/gateway/devices?connected=true",
			"macFiltering": "api/v1/service/macFiltering",
		},
	}
	router.RetrieveToken()
	return router
}
