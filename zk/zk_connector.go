package zk

import "github.com/stretchr/testify/mock"

// mockConnector is an autogenerated mock type for the connector type
type mockConnector struct {
	mock.Mock
}

// Connect provides a mock function with given fields: path
func (_m *mockConnector) Connect(path string) (connection, error) {
	ret := _m.Called(path)

	var r0 connection
	if rf, ok := ret.Get(0).(func(string) connection); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(connection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
