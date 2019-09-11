package index 

import (
	"net/http"
	"path/filepath"
	"strings"
	"st/pkg/el"
	"st/pkg/handlers/extra"
	"st/pkg/head"
	"st/pkg/server"
	"st/pkg/paths"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request) {
	path, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	if path == "/index/" {
		Main(s, w, r)
		return
	}

	if path == "/index/serial/" {
		Serial(s, w, r)
		return
	}

	if path == "/index/map.svg" {
		MapSVG(s, w, r)
		return
	}

	p := paths.Split(path)

	if p.Acronym == "" {
		h, err := s.Trees["index"].Search(p.Name, head.Lang(r.Host))
		if err != nil {
			s.Log.Println(err)
			http.NotFound(w, r)
			return
		}
		Hold(s, w, r, h)
		return
	}

	e, err := s.Trees["index"].LookupAcronym(p.Acronym)
	if err != nil {
		println(p.Name)
		path := strings.TrimRight(r.URL.Path, "/")
		http.Redirect(w, r, filepath.Dir(path) + "/", 301)
		return
	}

	h, ok := e.(*el.Hold)
	if ok {
		Hold(s, w, r, h)
		return
	}

	if p.Type != "" {
		extra.Files(s, w, r, p)
		return
	}
	/*

	serveIndexHoldOrEl(w, r, rp)
	*/
}


/*
func serveIndexHoldOrEl(w http.ResponseWriter, r *http.Request, p *paths.Path) {
	// Hold or El
	x, err := lookupAcronymMulti("index", rp.Acronym)
	if err != nil {
		debug(err)
		path := strings.TrimRight(r.URL.Path, "/")
		http.Redirect(w, r, filepath.Dir(path) + "/", 301)
		return
	}

	// Hold
	if treeType(x) == "hold" {
		serveIndexHold(w, r, x.(*el.Hold))
		return
	}

	// El
	http.Redirect(w, r, filepath.Dir(r.URL.Path) + "/", 301)
	return
	//serveIndexEl(w, r, rp.Page, rp.Acronym)
	//return
}

func treeType(x interface{}) string {
	switch x.(type) {
	case *el.Hold:
		return "hold"
	default:
		return "el"
	}
}

func serveIndexHoldNoAcronym(w http.ResponseWriter, r *http.Request, rp *reqPath) {
	h, err := srv.indexTree.Search(rp.Name, lang(r.Host))
	if err != nil {
		debug(err)
		http.NotFound(w, r)
		return
	}
	serveIndexHold(w, r, h)
}
*/
