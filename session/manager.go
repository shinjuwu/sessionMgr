package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"sync"
)

type Manager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxLifeTIme int64
	actor       Actor
}

var provides = make(map[string]Provider)

func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session:Register  called twice for provider " + name)
	}
	provides[name] = provider
}

func NewManager(provideName, cookieName string, maxLifeTIme int64) (*Manager, error) {
	if provider, ok := provides[provideName]; !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	} else {
		return &Manager{provider: provider, cookieName: cookieName, maxLifeTIme: maxLifeTIme}, nil
	}
}

func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
