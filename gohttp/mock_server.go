package gohttp

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
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

// StartMockServer starts the Mock server
func StartMockServer() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.enable = true
}

// StopMockServer stops the Mock server
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

// FlushMocks deletes all mocks
func FlushMocks() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.mocks = make(map[string]*Mock)
}

func (m *mockServer) key(method, url, body string) string {
	hasher := md5.New()
	hasher.Write([]byte(method + url + sanitize(body)))
	return hex.EncodeToString(hasher.Sum(nil))
}

func sanitize(s string) string {
	sanitized := strings.TrimSpace(s)
	if sanitized == "" {
		return sanitized
	}

	sanitized = strings.ReplaceAll(sanitized, "\t", "")
	sanitized = strings.ReplaceAll(sanitized, "\n", "")

	return sanitized
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
		Error: fmt.Errorf("no mock matching %s from %s", method, url),
	}
}
