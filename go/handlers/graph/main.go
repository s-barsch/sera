package graph

import (
	//"fmt"

	"fmt"
	"net/http"

	"g.rg-s.com/sera/go/server/meta"
	"g.rg-s.com/sera/go/viewer"
)

func MainRedirect(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := fmt.Sprintf("/%v/graph/2021/04", m.Lang)
		http.Redirect(w, r, path, http.StatusTemporaryRedirect)
	}
}
