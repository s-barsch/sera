package viewer

import (
	"errors"
	"net/http"

	"g.rg-s.com/sacer/go/requests/meta"
	"g.rg-s.com/sacer/go/server/engine"
	"g.rg-s.com/sacer/go/server/store"
	"github.com/sirupsen/logrus"
)

type Viewer struct {
	Logger *logrus.Logger
	Store  *store.Store
	Engine *engine.Engine
	Reload func() error
}

type HandleFunc func(v *Viewer, meta *meta.Meta) http.HandlerFunc

func New(logger *logrus.Logger, store *store.Store, engine *engine.Engine, reload func() error) (*Viewer, error) {
	if logger == nil {
		return nil, errors.New("logger cannot be nil")
	}
	if store == nil {
		return nil, errors.New("store cannot be nil")
	}
	if engine == nil {
		return nil, errors.New("engine cannot be nil")
	}
	if reload == nil {
		return nil, errors.New("reload function cannot be nil")
	}

	return &Viewer{
		Logger: logger,
		Store:  store,
		Engine: engine,
		Reload: reload,
	}, nil
}

func (v *Viewer) View(h func(v *Viewer, meta *meta.Meta) http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(v, &meta.Meta{})(w, r)
	}
}
