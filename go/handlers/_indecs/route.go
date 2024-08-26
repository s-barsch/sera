package indecs

import (
	"net/http"
	"path/filepath"

	"g.sacerb.com/sacer/go/handlers/extra"
	"g.sacerb.com/sacer/go/server"
	"g.sacerb.com/sacer/go/server/meta"
	"g.sacerb.com/sacer/go/server/paths"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {

	if !s.Flags.Local {
		http.Error(w, "temporarily unavailable", 503)
		return
	}

	reqPath, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	rel := reqPath[len("/indecs"):]

	if rel == "" {
		Main(s, w, r, m)
		return
	}

	if rel == "/serial" {
		Serial(s, w, r, m)
		return
	}

	if rel == "/map.svg" {
		MapIndex(s, w, r, m)
		return
	}

	if rel == "/map-all.svg" {
		MapAll(s, w, r, m)
		return
	}

	if rel == "/map.dot" {
		MapDot(s, w, r, m)
		return
	}

	p := paths.Split(reqPath)

	if p.IsFile() {
		extra.ServeFile(s, w, r, m, p)
		return
	}

	indecs := s.Trees["indecs"].Access(m.Auth.Subscriber)[m.Lang]

	if p.Hash == "" {
		t, err := indecs.SearchTree(p.Slug, m.Lang)
		if err != nil {
			s.Log.Println(err)
			http.NotFound(w, r)
			return
		}
		IndexPage(s, w, r, m, t)
		return
	}

	t, err := indecs.LookupTreeHash(p.Hash)
	if err != nil {
		http.Redirect(w, r, filepath.Dir(m.Path), 301)
		return
	}

	IndexPage(s, w, r, m, t)
}
