package index

import (
	"net/http"

	"g.sacerb.com/sacer/go/server"
	"g.sacerb.com/sacer/go/server/meta"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	Main(s, w, r, m)
}
