package viewer

import (
	"net/http"

	"g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
)

type Viewer struct {
	Store  *server.Store
	Engine *server.Engine
}

type handler func(v *Viewer, meta *meta.Meta) http.HandlerFunc

func (v *Viewer) View(h handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(v, &meta.Meta{})(w, r)
	}
}
