package index

import (
	"net/http"

	"g.rg-s.com/sferal/go/server"
	"g.rg-s.com/sferal/go/server/meta"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	Main(s, w, r, m)
}
