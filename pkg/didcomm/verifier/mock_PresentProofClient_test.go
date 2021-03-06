// Code generated by mockery v1.0.0. DO NOT EDIT.

package verifier

import (
	presentproof "github.com/hyperledger/aries-framework-go/pkg/client/presentproof"
	mock "github.com/stretchr/testify/mock"
)

// MockPresentProofClient is an autogenerated mock type for the PresentProofClient type
type MockPresentProofClient struct {
	mock.Mock
}

// SendRequestPresentation provides a mock function with given fields: msg, myDID, theirDID
func (_m *MockPresentProofClient) SendRequestPresentation(msg *presentproof.RequestPresentation, myDID string, theirDID string) (string, error) {
	ret := _m.Called(msg, myDID, theirDID)

	var r0 string
	if rf, ok := ret.Get(0).(func(*presentproof.RequestPresentation, string, string) string); ok {
		r0 = rf(msg, myDID, theirDID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*presentproof.RequestPresentation, string, string) error); ok {
		r1 = rf(msg, myDID, theirDID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
