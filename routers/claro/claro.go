package claro

import (
	"encoding/json"

	"github.com/alcmoraes/yip/routers"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

func (c *ClaroRouter) RefreshToken() error {
	resp, err := c.HTTPClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{
			"userName": c.Username,
			"password": c.Password,
		}).
		Post(c.Routes["login"])
	if err != nil {
		return err
	}

	var response RefreshTokenResponse
	json.Unmarshal(resp.Body(), &response)
	c.Token = response.AccessToken
	return nil
}

func (c *ClaroRouter) ListDevices() (output []routers.Device, err error) {
	resp, err := c.HTTPClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Access-Token", c.Token).
		Get(c.Routes["devices"])
	if err != nil {
		return output, err
	}

	json.Unmarshal(resp.Body(), &output)
	return output, nil
}

func (c *ClaroRouter) GetFilteredDevices() (output []routers.FilteredDevice, err error) {
	resp, err := c.HTTPClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Access-Token", c.Token).
		Get(c.Routes["macFiltering"])
	if err != nil {
		return output, err
	}

	var response FilteredDevicesResponse
	json.Unmarshal(resp.Body(), &response)
	return response.Rules, nil
}

func (c *ClaroRouter) FilterDeviceByMac(mac string) error {
	lockedDevices, err := c.GetFilteredDevices()
	if err != nil {
		return err
	}
	for _, device := range lockedDevices {
		if device.MacAddress == mac {
			return nil
		}
	}

	_, err = c.HTTPClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Access-Token", c.Token).
		SetBody(map[string]interface{}{
			"macAddress": mac,
			"active":     true,
		}).
		Post(c.Routes["macFiltering"])
	if err != nil {
		return err
	}

	return nil
}

func (c *ClaroRouter) UnfilterDeviceByMac(mac string) error {
	newRules := make([]FilterRule, 0)
	lockedDevices, err := c.GetFilteredDevices()
	if err != nil {
		return err
	}
	for _, device := range lockedDevices {
		if device.MacAddress != mac {
			newRules = append(newRules, FilterRule{
				Active:     true,
				MacAddress: device.MacAddress,
			})
		}
	}

	_, err = c.HTTPClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Access-Token", c.Token).
		SetBody(map[string]interface{}{
			"rules": newRules,
		}).
		Put(c.Routes["macFiltering"])
	if err != nil {
		return err
	}

	return nil
}

func (c *ClaroRouter) ClearMacFilters() error {
	_, err := c.HTTPClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Access-Token", c.Token).
		SetBody(map[string]interface{}{
			"rules": make([]string, 0),
		}).
		Put(c.Routes["macFiltering"])
	if err != nil {
		return err
	}
	return nil
}

func NewClaroRouter() routers.Router {
	router := &ClaroRouter{
		HTTPClient: resty.New().SetBaseURL(viper.GetString("claro.host")),
		Username:   viper.GetString("claro.username"),
		Password:   viper.GetString("claro.password"),
		Routes: map[string]string{
			"login":        "api/v1/gateway/users/login",
			"devices":      "api/v1/gateway/devices?connected=true",
			"macFiltering": "api/v1/service/macFiltering",
		},
	}
	if err := router.RefreshToken(); err != nil {
		panic(err)
	}
	return router
}
