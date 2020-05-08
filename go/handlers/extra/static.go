package extra

import (
	"fmt"
	"net/http"
	"path/filepath"
	"stferal/go/head"
	"stferal/go/paths"
	"stferal/go/server"
	"strings"
	"time"
)

func serveStatic(w http.ResponseWriter, r *http.Request, p string) {
	if filepath.Ext(p) == ".vtt" {
		w.Header().Set("Content-Type", "text/vtt")
	}
	w.Header().Set("Expires", time.Now().AddDate(0, 3, 0).Format(time.RFC1123))
	http.ServeFile(w, r, p)
}

func ServiceWorker(s *server.Server, w http.ResponseWriter, r *http.Request) {
	serveStatic(w, r, s.Paths.Data+"/static/js"+r.URL.Path)
}

func JSFiles(s *server.Server, w http.ResponseWriter, r *http.Request) {
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

func StaticFiles(s *server.Server, w http.ResponseWriter, r *http.Request) {
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

func RobotsFiles(s *server.Server, w http.ResponseWriter, r *http.Request) {
	path := fmt.Sprintf(
		"%v/static/seo/robots-%v.txt",
		s.Paths.Data,
		head.Lang(r.Host),
	)
	serveStatic(w, r, path)
}
