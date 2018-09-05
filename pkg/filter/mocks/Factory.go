// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import filter "github.com/olegsu/iris/pkg/filter"
import kube "github.com/olegsu/iris/pkg/kube"
import mock "github.com/stretchr/testify/mock"

// Factory is an autogenerated mock type for the Factory type
type Factory struct {
	mock.Mock
}

// Build provides a mock function with given fields: _a0, _a1, _a2
func (_m *Factory) Build(_a0 map[string]interface{}, _a1 filter.Service, _a2 kube.Kube) (filter.Filter, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 filter.Filter
	if rf, ok := ret.Get(0).(func(map[string]interface{}, filter.Service, kube.Kube) filter.Filter); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(filter.Filter)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]interface{}, filter.Service, kube.Kube) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}