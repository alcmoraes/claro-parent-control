package mock

import "github.com/alcmoraes/yip/routers"

type MockRouter struct {
}

func (m *MockRouter) RefreshToken() error {
	return nil
}
func (m *MockRouter) ListDevices() ([]routers.Device, error) {
	return []routers.Device{
		{
			MacAddress: "00:00:00:00:00:00",
			Name:       "Device 1",
		},
		{
			MacAddress: "0A:33:AB:CD:EF:00",
			Name:       "Device 2",
		},
	}, nil
}
func (m *MockRouter) GetFilteredDevices() ([]routers.FilteredDevice, error) {
	return []routers.FilteredDevice{
		{
			MacAddress: "00:00:00:00:00:00",
			Name:       "Device 1",
		},
	}, nil
}
func (m *MockRouter) FilterDeviceByMac(mac string) error {
	return nil
}
func (m *MockRouter) UnfilterDeviceByMac(mac string) error {
	return nil
}
func (m *MockRouter) ClearMacFilters() error {
	return nil
}

func NewMockRouter() routers.Router {
	return &MockRouter{}
}
