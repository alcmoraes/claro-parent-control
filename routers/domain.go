package routers

type FilteredDevice struct {
	Id         int    `json:"id"`         // : 0,
	Active     bool   `json:"active"`     // : true,
	Name       string `json:"name"`       // : "cpe-bridgefilter-37",
	MacAddress string `json:"macAddress"` // : "64:A2:00:01:E8:6B",
	DeviceId   int    `json:"deviceId"`   // : 0
}

type Device struct {
	Id         string `json:"id"`         // : "64:64:4A:39:E3:38",
	MacAddress string `json:"macAddress"` // : "64:64:4A:39:E3:38",
	Name       string `json:"name"`       // : "MiWiFi-R4CM",
	IpAddress  string `json:"ipAddress"`  // : "192.168.0.43",
}
