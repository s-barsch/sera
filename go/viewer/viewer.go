package viewer

import (
	"errors"
	"net/http"

	"g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
	"g.rg-s.com/sera/go/server/users"
	"github.com/sirupsen/logrus"
)

type Viewer struct {
	Logger *logrus.Logger
	Store  *server.Store
	Engine *server.Engine
	Users  *users.Users
	reload chan struct{}
}

type HandleFunc func(v *Viewer, meta *meta.Meta) http.HandlerFunc

func (v *Viewer) View(h func(v *Viewer, meta *meta.Meta) http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(v, &meta.Meta{})(w, r)
	}
}

func (v *Viewer) Reload() error {
	return errors.New("not implemented")
}
