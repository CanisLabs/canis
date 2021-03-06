// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	decorator "github.com/hyperledger/aries-framework-go/pkg/didcomm/protocol/decorator"

	mock "github.com/stretchr/testify/mock"

	schema "github.com/scoir/canis/pkg/schema"
)

// PresentationRegistry is an autogenerated mock type for the PresentationRegistry type
type PresentationRegistry struct {
	mock.Mock
}

// RequestPresentation provides a mock function with given fields: name, version, typ, attrInfo, predicateInfo
func (_m *PresentationRegistry) RequestPresentation(name string, version string, typ string, attrInfo map[string]*schema.IndyProofRequestAttr, predicateInfo map[string]*schema.IndyProofRequestPredicate) (*decorator.AttachmentData, error) {
	ret := _m.Called(name, version, typ, attrInfo, predicateInfo)

	var r0 *decorator.AttachmentData
	if rf, ok := ret.Get(0).(func(string, string, string, map[string]*schema.IndyProofRequestAttr, map[string]*schema.IndyProofRequestPredicate) *decorator.AttachmentData); ok {
		r0 = rf(name, version, typ, attrInfo, predicateInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*decorator.AttachmentData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, map[string]*schema.IndyProofRequestAttr, map[string]*schema.IndyProofRequestPredicate) error); ok {
		r1 = rf(name, version, typ, attrInfo, predicateInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Verify provides a mock function with given fields: format, presentation, request, theirDID, myDID
func (_m *PresentationRegistry) Verify(format string, presentation []byte, request []byte, theirDID string, myDID string) error {
	ret := _m.Called(format, presentation, request, theirDID, myDID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []byte, []byte, string, string) error); ok {
		r0 = rf(format, presentation, request, theirDID, myDID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
