package index

import (
	"net/http"
	"sacer/go/server"
	"sacer/go/server/meta"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	Main(s, w, r, m)
}
