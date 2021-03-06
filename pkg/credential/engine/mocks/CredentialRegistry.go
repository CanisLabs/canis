// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	decorator "github.com/hyperledger/aries-framework-go/pkg/didcomm/protocol/decorator"
	datastore "github.com/scoir/canis/pkg/datastore"

	mock "github.com/stretchr/testify/mock"
)

// CredentialRegistry is an autogenerated mock type for the CredentialRegistry type
type CredentialRegistry struct {
	mock.Mock
}

// CreateCredentialOffer provides a mock function with given fields: issuer, subjectDID, s, value
func (_m *CredentialRegistry) CreateCredentialOffer(issuer *datastore.DID, subjectDID string, s *datastore.Schema, value []byte) (string, *decorator.AttachmentData, error) {
	ret := _m.Called(issuer, subjectDID, s, value)

	var r0 string
	if rf, ok := ret.Get(0).(func(*datastore.DID, string, *datastore.Schema, []byte) string); ok {
		r0 = rf(issuer, subjectDID, s, value)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *decorator.AttachmentData
	if rf, ok := ret.Get(1).(func(*datastore.DID, string, *datastore.Schema, []byte) *decorator.AttachmentData); ok {
		r1 = rf(issuer, subjectDID, s, value)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*decorator.AttachmentData)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*datastore.DID, string, *datastore.Schema, []byte) error); ok {
		r2 = rf(issuer, subjectDID, s, value)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CreateSchema provides a mock function with given fields: s
func (_m *CredentialRegistry) CreateSchema(s *datastore.Schema) (string, error) {
	ret := _m.Called(s)

	var r0 string
	if rf, ok := ret.Get(0).(func(*datastore.Schema) string); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*datastore.Schema) error); ok {
		r1 = rf(s)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSchemaForProposal provides a mock function with given fields: format, data
func (_m *CredentialRegistry) GetSchemaForProposal(format string, data []byte) (string, error) {
	ret := _m.Called(format, data)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, []byte) string); ok {
		r0 = rf(format, data)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []byte) error); ok {
		r1 = rf(format, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IssueCredential provides a mock function with given fields: issuer, s, offerID, requestAttachment, values
func (_m *CredentialRegistry) IssueCredential(issuer *datastore.DID, s *datastore.Schema, offerID string, requestAttachment decorator.AttachmentData, values map[string]interface{}) (*decorator.AttachmentData, error) {
	ret := _m.Called(issuer, s, offerID, requestAttachment, values)

	var r0 *decorator.AttachmentData
	if rf, ok := ret.Get(0).(func(*datastore.DID, *datastore.Schema, string, decorator.AttachmentData, map[string]interface{}) *decorator.AttachmentData); ok {
		r0 = rf(issuer, s, offerID, requestAttachment, values)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*decorator.AttachmentData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*datastore.DID, *datastore.Schema, string, decorator.AttachmentData, map[string]interface{}) error); ok {
		r1 = rf(issuer, s, offerID, requestAttachment, values)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterSchema provides a mock function with given fields: registrant, s
func (_m *CredentialRegistry) RegisterSchema(registrant *datastore.DID, s *datastore.Schema) error {
	ret := _m.Called(registrant, s)

	var r0 error
	if rf, ok := ret.Get(0).(func(*datastore.DID, *datastore.Schema) error); ok {
		r0 = rf(registrant, s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
