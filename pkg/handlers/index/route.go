package index

import (
	"net/http"
	"path/filepath"
	"stferal/pkg/entry"
	"stferal/pkg/handlers/extra"
	"stferal/pkg/head"
	"stferal/pkg/paths"
	"stferal/pkg/server"
	"strings"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request) {
	path, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	rel := path[len("/index"):]

	if rel == "" {
		Main(s, w, r)
		return
	}

	if rel == "/serial/" {
		Serial(s, w, r)
		return
	}

	if rel == "/map.svg" {
		MapSVG(s, w, r)
		return
	}

	p := paths.Split(path)

	if p.Type != "" {
		extra.Files(s, w, r, p)
		return
	}

	tree := s.Trees["index"]

	if s.Flags.Local {
		tree = s.Trees["index-private"]
	}

	if p.Acronym == "" {
		h, err := tree.Search(p.Name, head.Lang(r.Host))
		if err != nil {
			s.Log.Println(err)
			http.NotFound(w, r)
			return
		}
		Hold(s, w, r, h)
		return
	}

	e, err := tree.LookupAcronym(p.Acronym)
	if err != nil {
		println("hare")
		path := strings.TrimRight(r.URL.Path, "/")
		http.Redirect(w, r, filepath.Dir(path)+"/", 301)
		return
	}

	h, ok := e.(*entry.Hold)
	if ok {
		Hold(s, w, r, h)
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
		serveIndexHold(w, r, x.(*entry.Hold))
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
	case *entry.Hold:
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
