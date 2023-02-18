package routers

type Router interface {
	RefreshToken() error
	ListDevices() ([]Device, error)
	GetFilteredDevices() ([]FilteredDevice, error)
	FilterDeviceByMac(mac string) error
	UnfilterDeviceByMac(mac string) error
	ClearMacFilters() error
}
