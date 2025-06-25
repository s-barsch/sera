package extra

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	s "g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
	"g.rg-s.com/sera/go/server/paths"
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

func ServiceWorker(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	serveStatic(w, r, s.Srv.Paths.Data+"/static/js"+r.URL.Path)
}

func JSFiles(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	path, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		s.Srv.Debug(err)
		http.NotFound(w, r)
		return
	}

	// Block folders.
	if strings.HasSuffix(path, "/") {
		http.NotFound(w, r)
		return
	}

	serveStatic(w, r, s.Srv.Paths.Data+"/static"+path)
}

func StaticFiles(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	path, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		s.Srv.Debug(err)
		http.NotFound(w, r)
		return
	}

	// Block folders.
	if strings.HasSuffix(path, "/") {
		http.NotFound(w, r)
		return
	}

	serveStatic(w, r, s.Srv.Paths.Data+path)
}

func RobotsFiles(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	path := fmt.Sprintf("%v/static/seo/robots-de.txt", s.Srv.Paths.Data)
	serveStatic(w, r, path)
}
