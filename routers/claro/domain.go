package claro

import "github.com/go-resty/resty/v2"

type RetrieveTokenResponse struct {
	AccessToken string `json:"accessToken"`
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