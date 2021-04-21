// Code generated by mockery v1.0.0. DO NOT EDIT.

package issuecredential

import (
	service "github.com/hyperledger/aries-framework-go/pkg/didcomm/common/service"
	"github.com/hyperledger/aries-framework-go/pkg/didcomm/protocol/issuecredential"
	mock "github.com/stretchr/testify/mock"
)

// MockMetadata is an autogenerated mock type for the Metadata type
type MockMetadata struct {
	mock.Mock
}

// CredentialNames provides a mock function with given fields:
func (_m *MockMetadata) CredentialNames() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// IssueCredential provides a mock function with given fields:
func (_m *MockMetadata) IssueCredential() *issuecredential.IssueCredential {
	ret := _m.Called()

	var r0 *issuecredential.IssueCredential
	if rf, ok := ret.Get(0).(func() *issuecredential.IssueCredential); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*issuecredential.IssueCredential)
		}
	}

	return r0
}

// Message provides a mock function with given fields:
func (_m *MockMetadata) Message() service.DIDCommMsg {
	ret := _m.Called()

	var r0 service.DIDCommMsg
	if rf, ok := ret.Get(0).(func() service.DIDCommMsg); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(service.DIDCommMsg)
		}
	}

	return r0
}

// OfferCredential provides a mock function with given fields:
func (_m *MockMetadata) OfferCredential() *issuecredential.OfferCredential {
	ret := _m.Called()

	var r0 *issuecredential.OfferCredential
	if rf, ok := ret.Get(0).(func() *issuecredential.OfferCredential); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*issuecredential.OfferCredential)
		}
	}

	return r0
}

// Properties provides a mock function with given fields:
func (_m *MockMetadata) Properties() map[string]interface{} {
	ret := _m.Called()

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func() map[string]interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// ProposeCredential provides a mock function with given fields:
func (_m *MockMetadata) ProposeCredential() *issuecredential.ProposeCredential {
	ret := _m.Called()

	var r0 *issuecredential.ProposeCredential
	if rf, ok := ret.Get(0).(func() *issuecredential.ProposeCredential); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*issuecredential.ProposeCredential)
		}
	}

	return r0
}

// StateName provides a mock function with given fields:
func (_m *MockMetadata) StateName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}