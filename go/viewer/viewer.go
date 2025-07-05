package viewer

import (
	"errors"
	"net/http"

	"g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
	"g.rg-s.com/sera/go/server/users"
	"g.rg-s.com/sera/go/store"
	"github.com/sirupsen/logrus"
)

type Viewer struct {
	Logger *logrus.Logger
	Store  *store.Store
	Engine *server.Engine
	Users  *users.Users
	Reload func() error
}

type HandleFunc func(v *Viewer, meta *meta.Meta) http.HandlerFunc

func NewViewer(logger *logrus.Logger, store *store.Store, engine *server.Engine, users *users.Users, reload func() error) (*Viewer, error) {
	if logger == nil {
		return nil, errors.New("logger cannot be nil")
	}
	if store == nil {
		return nil, errors.New("store cannot be nil")
	}
	if engine == nil {
		return nil, errors.New("engine cannot be nil")
	}

	return &Viewer{
		Logger: logger,
		Store:  store,
		Engine: engine,
		Users:  users,
		Reload: reload,
	}, nil
}

func (v *Viewer) View(h func(v *Viewer, meta *meta.Meta) http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(v, &meta.Meta{})(w, r)
	}
}
