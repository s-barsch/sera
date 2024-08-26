package extra

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"g.sacerb.com/sacer/go/server"
	"g.sacerb.com/sacer/go/server/meta"
	"g.sacerb.com/sacer/go/server/paths"
)

func serveStatic(w http.ResponseWriter, r *http.Request, p string) {
	switch filepath.Ext(p) {
	case ".vtt":
		w.Header().Set("Content-Type", "text/vtt")
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".js":
		w.Header().Set("Content-Type", "text/javascript")
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
	path := fmt.Sprintf("%v/static/seo/robots-de.txt", s.Paths.Data)
	serveStatic(w, r, path)
}
