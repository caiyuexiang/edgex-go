// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/edgexfoundry/go-mod-core-contracts/v2/models"

// DeviceLoader is an autogenerated mock type for the DeviceLoader type
type DeviceLoader struct {
	mock.Mock
}

// GetAllDevices provides a mock function with given fields:
func (_m *DeviceLoader) GetAllDevices() ([]models.Device, error) {
	ret := _m.Called()

	var r0 []models.Device
	if rf, ok := ret.Get(0).(func() []models.Device); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDevicesByProfileId provides a mock function with given fields: pid
func (_m *DeviceLoader) GetDevicesByProfileId(pid string) ([]models.Device, error) {
	ret := _m.Called(pid)

	var r0 []models.Device
	if rf, ok := ret.Get(0).(func(string) []models.Device); ok {
		r0 = rf(pid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(pid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
