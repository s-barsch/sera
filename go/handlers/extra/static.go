package extra

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"g.rg-s.com/sacer/go/requests/meta"
	"g.rg-s.com/sacer/go/requests/split"
	"g.rg-s.com/sacer/go/viewer"
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

func ServiceWorker(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		serveStatic(w, r, v.Engine.Config.Paths.Data+"/static/js"+r.URL.Path)
	}
}

func JSFiles(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path, err := split.Sanitize(r.URL.Path)
		if err != nil {
			v.Logger.Info(err)
			http.NotFound(w, r)
			return
		}

		// Block folders.
		if strings.HasSuffix(path, "/") {
			http.NotFound(w, r)
			return
		}

		serveStatic(w, r, v.Engine.Config.Paths.Data+"/static"+path)
	}
}

func StaticFiles(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path, err := split.Sanitize(r.URL.Path)
		if err != nil {
			v.Logger.Info(err)
			http.NotFound(w, r)
			return
		}

		// Block folders.
		if strings.HasSuffix(path, "/") {
			http.NotFound(w, r)
			return
		}

		serveStatic(w, r, v.Engine.Config.Paths.Data+path)
	}
}

func RobotsFiles(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := fmt.Sprintf("%v/static/seo/robots-de.txt", v.Engine.Config.Paths.Data)
		serveStatic(w, r, path)
	}
}
