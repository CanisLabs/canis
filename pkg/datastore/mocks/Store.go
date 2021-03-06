// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	didexchange "github.com/hyperledger/aries-framework-go/pkg/client/didexchange"
	datastore "github.com/scoir/canis/pkg/datastore"
	mock "github.com/stretchr/testify/mock"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// AddWebhook provides a mock function with given fields: hook
func (_m *Store) AddWebhook(hook *datastore.Webhook) error {
	ret := _m.Called(hook)

	var r0 error
	if rf, ok := ret.Get(0).(func(*datastore.Webhook) error); ok {
		r0 = rf(hook)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAgent provides a mock function with given fields: name
func (_m *Store) DeleteAgent(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAgentConnection provides a mock function with given fields: a, externalID
func (_m *Store) DeleteAgentConnection(a *datastore.Agent, externalID string) error {
	ret := _m.Called(a, externalID)

	var r0 error
	if rf, ok := ret.Get(0).(func(*datastore.Agent, string) error); ok {
		r0 = rf(a, externalID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCloudAgentConnection provides a mock function with given fields: a, externalID
func (_m *Store) DeleteCloudAgentConnection(a *datastore.CloudAgent, externalID string) error {
	ret := _m.Called(a, externalID)

	var r0 error
	if rf, ok := ret.Get(0).(func(*datastore.CloudAgent, string) error); ok {
		r0 = rf(a, externalID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCloudAgentCredential provides a mock function with given fields: a, id
func (_m *Store) DeleteCloudAgentCredential(a *datastore.CloudAgent, id string) error {
	ret := _m.Called(a, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(*datastore.CloudAgent, string) error); ok {
		r0 = rf(a, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCredentialByOffer provides a mock function with given fields: offerID
func (_m *Store) DeleteCredentialByOffer(offerID string) error {
	ret := _m.Called(offerID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(offerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSchema provides a mock function with given fields: name
func (_m *Store) DeleteSchema(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteWebhook provides a mock function with given fields: typ
func (_m *Store) DeleteWebhook(typ string) error {
	ret := _m.Called(typ)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(typ)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindCredentialByProtocolID provides a mock function with given fields: offerID
func (_m *Store) FindCredentialByProtocolID(offerID string) (*datastore.IssuedCredential, error) {
	ret := _m.Called(offerID)

	var r0 *datastore.IssuedCredential
	if rf, ok := ret.Get(0).(func(string) *datastore.IssuedCredential); ok {
		r0 = rf(offerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datastore.IssuedCredential)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(offerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAgent provides a mock function with given fields: id
func (_m *Store) GetAgent(id string) (*datastore.Agent, error) {
	ret := _m.Called(id)

	var r0 *datastore.Agent
	if rf, ok := ret.Get(0).(func(string) *datastore.Agent); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datastore.Agent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAgentByPublicDID provides a mock function with given fields: DID
func (_m *Store) GetAgentByPublicDID(DID string) (*datastore.Agent, error) {
	ret := _m.Called(DID)

	var r0 *datastore.Agent
	if rf, ok := ret.Get(0).(func(string) *datastore.Agent); ok {
		r0 = rf(DID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datastore.Agent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(DID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAgentConnection provides a mock function with given fields: a, externalID
func (_m *Store) GetAgentConnection(a *datastore.Agent, externalID string) (*datastore.AgentConnection, error) {
	ret := _m.Called(a, externalID)

	var r0 *datastore.AgentConnection
	if rf, ok := ret.Get(0).(func(*datastore.Agent, string) *datastore.AgentConnection); ok {
		r0 = rf(a, externalID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datastore.AgentConnection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*datastore.Agent, string) error); ok {
		r1 = rf(a, externalID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAgentConnectionForDID provides a mock function with given fields: a, theirDID
func (_m *Store) GetAgentConnectionForDID(a *datastore.Agent, theirDID string) (*datastore.AgentConnection, error) {
	ret := _m.Called(a, theirDID)

	var r0 *datastore.AgentConnection
	if rf, ok := ret.Get(0).(func(*datastore.Agent, string) *datastore.AgentConnection); ok {
		r0 = rf(a, theirDID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datastore.AgentConnection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*datastore.Agent, string) error); ok {
		r1 = rf(a, theirDID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCloudAgent provides a mock function with given fields: ID
func (_m *Store) GetCloudAgent(ID string) (*datastore.CloudAgent, error) {
	ret := _m.Called(ID)

	var r0 *datastore.CloudAgent
	if rf, ok := ret.Get(0).(func(string) *datastore.CloudAgent); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datastore.CloudAgent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCloudAgentConnection provides a mock function with given fields: a, invitationID
func (_m *Store) GetCloudAgentConnection(a *datastore.CloudAgent, invitationID string) (*datastore.CloudAgentConnection, error) {
	ret := _m.Called(a, invitationID)

	var r0 *datastore.CloudAgentConnection
	if rf, ok := ret.Get(0).(func(*datastore.CloudAgent, string) *datastore.CloudAgentConnection); ok {
		r0 = rf(a, invitationID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datastore.CloudAgentConnection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*datastore.CloudAgent, string) error); ok {
		r1 = rf(a, invitationID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCloudAgentConnectionForDIDs provides a mock function with given fields: myDID, theirDID
func (_m *Store) GetCloudAgentConnectionForDIDs(myDID string, theirDID string) (*datastore.CloudAgentConnection, error) {
	ret := _m.Called(myDID, theirDID)

	var r0 *datastore.CloudAgentConnection
	if rf, ok := ret.Get(0).(func(string, string) *datastore.CloudAgentConnection); ok {
		r0 = rf(myDID, theirDID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datastore.CloudAgentConnection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(myDID, theirDID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCloudAgentCredential provides a mock function with given fields: a, id
func (_m *Store) GetCloudAgentCredential(a *datastore.CloudAgent, id string) (*datastore.CloudAgentCredential, error) {
	ret := _m.Called(a, id)

	var r0 *datastore.CloudAgentCredential
	if rf, ok := ret.Get(0).(func(*datastore.CloudAgent, string) *datastore.CloudAgentCredential); ok {
		r0 = rf(a, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datastore.CloudAgentCredential)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*datastore.CloudAgent, string) error); ok {
		r1 = rf(a, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCloudAgentCredentialFromThread provides a mock function with given fields: cloudAgentID, thid
func (_m *Store) GetCloudAgentCredentialFromThread(cloudAgentID string, thid string) (*datastore.CloudAgentCredential, error) {
	ret := _m.Called(cloudAgentID, thid)

	var r0 *datastore.CloudAgentCredential
	if rf, ok := ret.Get(0).(func(string, string) *datastore.CloudAgentCredential); ok {
		r0 = rf(cloudAgentID, thid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datastore.CloudAgentCredential)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(cloudAgentID, thid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCloudAgentForDID provides a mock function with given fields: theirDID
func (_m *Store) GetCloudAgentForDID(theirDID string) (*datastore.CloudAgent, error) {
	ret := _m.Called(theirDID)

	var r0 *datastore.CloudAgent
	if rf, ok := ret.Get(0).(func(string) *datastore.CloudAgent); ok {
		r0 = rf(theirDID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datastore.CloudAgent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(theirDID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDID provides a mock function with given fields: id
func (_m *Store) GetDID(id string) (*datastore.DID, error) {
	ret := _m.Called(id)

	var r0 *datastore.DID
	if rf, ok := ret.Get(0).(func(string) *datastore.DID); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datastore.DID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEdgeAgent provides a mock function with given fields: connectionID
func (_m *Store) GetEdgeAgent(connectionID string) (*datastore.EdgeAgent, error) {
	ret := _m.Called(connectionID)

	var r0 *datastore.EdgeAgent
	if rf, ok := ret.Get(0).(func(string) *datastore.EdgeAgent); ok {
		r0 = rf(connectionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datastore.EdgeAgent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(connectionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEdgeAgentForDID provides a mock function with given fields: theirDID
func (_m *Store) GetEdgeAgentForDID(theirDID string) (*datastore.EdgeAgent, error) {
	ret := _m.Called(theirDID)

	var r0 *datastore.EdgeAgent
	if rf, ok := ret.Get(0).(func(string) *datastore.EdgeAgent); ok {
		r0 = rf(theirDID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datastore.EdgeAgent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(theirDID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMediatorDID provides a mock function with given fields:
func (_m *Store) GetMediatorDID() (*datastore.DID, error) {
	ret := _m.Called()

	var r0 *datastore.DID
	if rf, ok := ret.Get(0).(func() *datastore.DID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datastore.DID)
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

// GetPresentationRequest provides a mock function with given fields: ID
func (_m *Store) GetPresentationRequest(ID string) (*datastore.PresentationRequest, error) {
	ret := _m.Called(ID)

	var r0 *datastore.PresentationRequest
	if rf, ok := ret.Get(0).(func(string) *datastore.PresentationRequest); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datastore.PresentationRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPublicDID provides a mock function with given fields:
func (_m *Store) GetPublicDID() (*datastore.DID, error) {
	ret := _m.Called()

	var r0 *datastore.DID
	if rf, ok := ret.Get(0).(func() *datastore.DID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datastore.DID)
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

// GetSchema provides a mock function with given fields: name
func (_m *Store) GetSchema(name string) (*datastore.Schema, error) {
	ret := _m.Called(name)

	var r0 *datastore.Schema
	if rf, ok := ret.Get(0).(func(string) *datastore.Schema); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datastore.Schema)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSchemaByExternalID provides a mock function with given fields: externalID
func (_m *Store) GetSchemaByExternalID(externalID string) (*datastore.Schema, error) {
	ret := _m.Called(externalID)

	var r0 *datastore.Schema
	if rf, ok := ret.Get(0).(func(string) *datastore.Schema); ok {
		r0 = rf(externalID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datastore.Schema)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(externalID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertAgent provides a mock function with given fields: a
func (_m *Store) InsertAgent(a *datastore.Agent) (string, error) {
	ret := _m.Called(a)

	var r0 string
	if rf, ok := ret.Get(0).(func(*datastore.Agent) string); ok {
		r0 = rf(a)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*datastore.Agent) error); ok {
		r1 = rf(a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertAgentConnection provides a mock function with given fields: a, externalID, conn
func (_m *Store) InsertAgentConnection(a *datastore.Agent, externalID string, conn *didexchange.Connection) error {
	ret := _m.Called(a, externalID, conn)

	var r0 error
	if rf, ok := ret.Get(0).(func(*datastore.Agent, string, *didexchange.Connection) error); ok {
		r0 = rf(a, externalID, conn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertCloudAgentConnection provides a mock function with given fields: ac
func (_m *Store) InsertCloudAgentConnection(ac *datastore.CloudAgentConnection) error {
	ret := _m.Called(ac)

	var r0 error
	if rf, ok := ret.Get(0).(func(*datastore.CloudAgentConnection) error); ok {
		r0 = rf(ac)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertCloudAgentCredential provides a mock function with given fields: cred
func (_m *Store) InsertCloudAgentCredential(cred *datastore.CloudAgentCredential) error {
	ret := _m.Called(cred)

	var r0 error
	if rf, ok := ret.Get(0).(func(*datastore.CloudAgentCredential) error); ok {
		r0 = rf(cred)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertCredential provides a mock function with given fields: c
func (_m *Store) InsertCredential(c *datastore.IssuedCredential) (string, error) {
	ret := _m.Called(c)

	var r0 string
	if rf, ok := ret.Get(0).(func(*datastore.IssuedCredential) string); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*datastore.IssuedCredential) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertDID provides a mock function with given fields: d
func (_m *Store) InsertDID(d *datastore.DID) error {
	ret := _m.Called(d)

	var r0 error
	if rf, ok := ret.Get(0).(func(*datastore.DID) error); ok {
		r0 = rf(d)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertPresentation provides a mock function with given fields: p
func (_m *Store) InsertPresentation(p *datastore.Presentation) (string, error) {
	ret := _m.Called(p)

	var r0 string
	if rf, ok := ret.Get(0).(func(*datastore.Presentation) string); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*datastore.Presentation) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertPresentationRequest provides a mock function with given fields: pr
func (_m *Store) InsertPresentationRequest(pr *datastore.PresentationRequest) (string, error) {
	ret := _m.Called(pr)

	var r0 string
	if rf, ok := ret.Get(0).(func(*datastore.PresentationRequest) string); ok {
		r0 = rf(pr)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*datastore.PresentationRequest) error); ok {
		r1 = rf(pr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertSchema provides a mock function with given fields: s
func (_m *Store) InsertSchema(s *datastore.Schema) (string, error) {
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

// ListAgent provides a mock function with given fields: c
func (_m *Store) ListAgent(c *datastore.AgentCriteria) (*datastore.AgentList, error) {
	ret := _m.Called(c)

	var r0 *datastore.AgentList
	if rf, ok := ret.Get(0).(func(*datastore.AgentCriteria) *datastore.AgentList); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datastore.AgentList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*datastore.AgentCriteria) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAgentConnections provides a mock function with given fields: a
func (_m *Store) ListAgentConnections(a *datastore.Agent) ([]*datastore.AgentConnection, error) {
	ret := _m.Called(a)

	var r0 []*datastore.AgentConnection
	if rf, ok := ret.Get(0).(func(*datastore.Agent) []*datastore.AgentConnection); ok {
		r0 = rf(a)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*datastore.AgentConnection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*datastore.Agent) error); ok {
		r1 = rf(a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCloudAgentConnections provides a mock function with given fields: a
func (_m *Store) ListCloudAgentConnections(a *datastore.CloudAgent) ([]*datastore.CloudAgentConnection, error) {
	ret := _m.Called(a)

	var r0 []*datastore.CloudAgentConnection
	if rf, ok := ret.Get(0).(func(*datastore.CloudAgent) []*datastore.CloudAgentConnection); ok {
		r0 = rf(a)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*datastore.CloudAgentConnection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*datastore.CloudAgent) error); ok {
		r1 = rf(a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCloudAgentCredentials provides a mock function with given fields: a
func (_m *Store) ListCloudAgentCredentials(a *datastore.CloudAgent) ([]*datastore.CloudAgentCredential, error) {
	ret := _m.Called(a)

	var r0 []*datastore.CloudAgentCredential
	if rf, ok := ret.Get(0).(func(*datastore.CloudAgent) []*datastore.CloudAgentCredential); ok {
		r0 = rf(a)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*datastore.CloudAgentCredential)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*datastore.CloudAgent) error); ok {
		r1 = rf(a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDIDs provides a mock function with given fields: c
func (_m *Store) ListDIDs(c *datastore.DIDCriteria) (*datastore.DIDList, error) {
	ret := _m.Called(c)

	var r0 *datastore.DIDList
	if rf, ok := ret.Get(0).(func(*datastore.DIDCriteria) *datastore.DIDList); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datastore.DIDList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*datastore.DIDCriteria) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSchema provides a mock function with given fields: c
func (_m *Store) ListSchema(c *datastore.SchemaCriteria) (*datastore.SchemaList, error) {
	ret := _m.Called(c)

	var r0 *datastore.SchemaList
	if rf, ok := ret.Get(0).(func(*datastore.SchemaCriteria) *datastore.SchemaList); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datastore.SchemaList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*datastore.SchemaCriteria) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListWebhooks provides a mock function with given fields: typ
func (_m *Store) ListWebhooks(typ string) ([]*datastore.Webhook, error) {
	ret := _m.Called(typ)

	var r0 []*datastore.Webhook
	if rf, ok := ret.Get(0).(func(string) []*datastore.Webhook); ok {
		r0 = rf(typ)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*datastore.Webhook)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(typ)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterCloudAgent provides a mock function with given fields: externalID, publicKey, nextKey
func (_m *Store) RegisterCloudAgent(externalID string, publicKey []byte, nextKey []byte) (string, error) {
	ret := _m.Called(externalID, publicKey, nextKey)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, []byte, []byte) string); ok {
		r0 = rf(externalID, publicKey, nextKey)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []byte, []byte) error); ok {
		r1 = rf(externalID, publicKey, nextKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterEdgeAgent provides a mock function with given fields: connectionID, externalID
func (_m *Store) RegisterEdgeAgent(connectionID string, externalID string) (string, error) {
	ret := _m.Called(connectionID, externalID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(connectionID, externalID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(connectionID, externalID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetMediatorDID provides a mock function with given fields: DID
func (_m *Store) SetMediatorDID(DID *datastore.DID) error {
	ret := _m.Called(DID)

	var r0 error
	if rf, ok := ret.Get(0).(func(*datastore.DID) error); ok {
		r0 = rf(DID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetPublicDID provides a mock function with given fields: DID
func (_m *Store) SetPublicDID(DID *datastore.DID) error {
	ret := _m.Called(DID)

	var r0 error
	if rf, ok := ret.Get(0).(func(*datastore.DID) error); ok {
		r0 = rf(DID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAgent provides a mock function with given fields: a
func (_m *Store) UpdateAgent(a *datastore.Agent) error {
	ret := _m.Called(a)

	var r0 error
	if rf, ok := ret.Get(0).(func(*datastore.Agent) error); ok {
		r0 = rf(a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateCloudAgent provides a mock function with given fields: ea
func (_m *Store) UpdateCloudAgent(ea *datastore.CloudAgent) error {
	ret := _m.Called(ea)

	var r0 error
	if rf, ok := ret.Get(0).(func(*datastore.CloudAgent) error); ok {
		r0 = rf(ea)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateCloudAgentConnection provides a mock function with given fields: ac
func (_m *Store) UpdateCloudAgentConnection(ac *datastore.CloudAgentConnection) error {
	ret := _m.Called(ac)

	var r0 error
	if rf, ok := ret.Get(0).(func(*datastore.CloudAgentConnection) error); ok {
		r0 = rf(ac)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateCloudAgentCredential provides a mock function with given fields: cred
func (_m *Store) UpdateCloudAgentCredential(cred *datastore.CloudAgentCredential) error {
	ret := _m.Called(cred)

	var r0 error
	if rf, ok := ret.Get(0).(func(*datastore.CloudAgentCredential) error); ok {
		r0 = rf(cred)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateCredential provides a mock function with given fields: c
func (_m *Store) UpdateCredential(c *datastore.IssuedCredential) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(*datastore.IssuedCredential) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateEdgeAgent provides a mock function with given fields: ea
func (_m *Store) UpdateEdgeAgent(ea *datastore.EdgeAgent) error {
	ret := _m.Called(ea)

	var r0 error
	if rf, ok := ret.Get(0).(func(*datastore.EdgeAgent) error); ok {
		r0 = rf(ea)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateSchema provides a mock function with given fields: s
func (_m *Store) UpdateSchema(s *datastore.Schema) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(*datastore.Schema) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
