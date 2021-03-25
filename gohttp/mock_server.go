package gohttp

import (
	"errors"
	"fmt"
	"sync"
)

var (
	mockupServer = mockServer{
		mocks: make(map[string]*Mock),
	}
)

type mockServer struct {
	enable      bool
	serverMutex sync.Mutex
	mocks       map[string]*Mock
}

// Starts the Mock server
func StartMockServer() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.enable = true
}

// Stops the Mock server
func StopMockServer() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.enable = false
}

// AddMock adds a mock to the collection of mocks
func AddMock(mock *Mock) {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	k := mockupServer.key(mock.Method, mock.URL, mock.RequestBody)
	mockupServer.mocks[k] = mock
}

func (m *mockServer) key(method, url, body string) string {
	return method + url + body
}

func (m *mockServer) mock(method, url, body string) *Mock {
	if !m.enable {
		return nil
	}

	mock := m.mocks[m.key(method, url, body)]
	if mock != nil {
		return mock
	}

	return &Mock{
		Error: errors.New(fmt.Sprintf("no mock matching %s from %s", method, url)),
	}
}
