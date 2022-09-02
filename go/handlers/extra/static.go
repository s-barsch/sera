package extra

import (
	"fmt"
	"net/http"
	"path/filepath"
	"sacer/go/server"
	"sacer/go/server/meta"
	"sacer/go/server/paths"
	"strings"
	"time"
)

func serveStatic(w http.ResponseWriter, r *http.Request, p string) {
	switch filepath.Ext(p) {
	case ".vtt":
		w.Header().Set("Content-Type", "text/vtt")
	}
	w.Header().Set("Expires", time.Now().AddDate(0, 3, 0).Format(time.RFC1123))
	http.ServeFile(w, r, p)
}

func ServiceWorker(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	serveStatic(w, r, s.Paths.Data+"/static/js"+r.URL.Path)
}

func JSFiles(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	path, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		s.Debug(err)
		http.NotFound(w, r)
		return
	}

	// Block folders.
	if strings.HasSuffix(path, "/") {
		http.NotFound(w, r)
		return
	}

	serveStatic(w, r, s.Paths.Data+"/static"+path)
}

func StaticFiles(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	path, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		s.Debug(err)
		http.NotFound(w, r)
		return
	}

	// Block folders.
	if strings.HasSuffix(path, "/") {
		http.NotFound(w, r)
		return
	}

	serveStatic(w, r, s.Paths.Data+path)
}

func RobotsFiles(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	path := fmt.Sprintf(
		"%v/static/seo/robots-%v.txt",
		s.Paths.Data,
		m.Lang,
	)
	serveStatic(w, r, path)
}
