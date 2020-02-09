package sitemaps

import (
	"fmt"
	"log"
	"net/http"
	"stferal/pkg/entry"
	"stferal/pkg/head"
	"stferal/pkg/server"
	"time"
)

type Entry struct {
	Loc      string
	Lastmod  string
	Priority string
}

func Index(s *server.Server, w http.ResponseWriter, r *http.Request) {
	domain := "https://stferal.com"
	if head.Lang(r.Host) == "en" {
		domain = "https://en.stferal.com"
	}
	err := s.Templates.ExecuteTemplate(w, "sitemap-index", struct{ Domain string }{domain})
	if err != nil {
		log.Println(err)
		return
	}
}

func Core(s *server.Server, w http.ResponseWriter, r *http.Request) {
	entries, err := coreEntries(s, head.Lang(r.Host))
	if err != nil {
		http.Error(w, "internal error", 500)
		log.Println(err)
		return
	}

	err = s.Templates.ExecuteTemplate(w, "sitemap", entries)
	if err != nil {
		log.Println(err)
		return
	}
}

func Holds(s *server.Server, w http.ResponseWriter, r *http.Request) {
	entries := categoryEntries(s, head.Lang(r.Host))

	entries = append(entries, holdEntries(s, head.Lang(r.Host))...)

	err := s.Templates.ExecuteTemplate(w, "sitemap", entries)
	if err != nil {
		log.Println(err)
		return
	}
}

/*
func IndexEls(w http.ResponseWriter, r *http.Request) {
	entries, err := smEls("index", lang(r.Host))
	if err != nil {
		http.Error(w, "internal error", 500)
		log.Println(err)
		return
	}

	err = srv.tmpls.ExecuteTemplate(w, "sitemap", entries)
	if err != nil {
		log.Println(err)
		return
	}
}
*/

func GraphEls(s *server.Server, w http.ResponseWriter, r *http.Request) {
	entries, err := elEntries(s, "graph", head.Lang(r.Host))
	if err != nil {
		http.Error(w, "internal error", 500)
		log.Println(err)
		return
	}

	err = s.Templates.ExecuteTemplate(w, "sitemap", entries)
	if err != nil {
		log.Println(err)
		return
	}
}

func coreEntries(s *server.Server, lang string) ([]*Entry, error) {
	entries := []*Entry{}

	tIndex, err := entry.DateSafe(s.Recents["index"][0])
	if err != nil {
		return nil, err
	}

	tGraph, err := entry.DateSafe(s.Recents["graph"][0])
	if err != nil {
		return nil, err
	}

	for _, v := range head.NewNav(lang) {
		priority := "0.9"
		lastmod := time.Time{}

		switch v.Name {
		case "home":
			priority = "1.0"
			if tIndex.Unix() > tGraph.Unix() {
				lastmod = tIndex
			} else {
				lastmod = tGraph
			}
		case "index":
			lastmod = tIndex
		case "graph":
			lastmod = tGraph
		case "über", "about":
			lastmod = s.Trees["about"].File.ModTime
		}

		entries = append(entries, &Entry{
			Loc:      absoluteURL(v.Href, lang),
			Priority: priority,
			Lastmod:  lastmod.Format(time.RFC3339),
		})
	}
	return entries, nil
}

func categoryEntries(s *server.Server, lang string) []*Entry {
	entries := []*Entry{}
	trees := []*entry.Hold{
		s.Trees["graph"],
		s.Trees["index"],
	}
	for _, t := range trees {
		for _, h := range t.Holds {
			if h.Depth() != 1 {
				continue
			}
			entries = append(entries, &Entry{
				Loc:      absoluteURL(h.Permalink(lang), lang),
				Lastmod:  h.File.ModTime.Format(time.RFC3339),
				Priority: "0.7",
			})
		}
	}
	return entries
}

func holdEntries(s *server.Server, lang string) []*Entry {
	return append(indexHolds(s, lang), aboutHolds(s, lang)...)
}

func aboutHolds(s *server.Server, lang string) []*Entry {
	entries := []*Entry{}
	holds := s.Trees["about"].TraverseHolds()
	for _, h := range holds {
		entries = append(entries, &Entry{
			Loc:      absoluteURL(h.Permalink(lang), lang),
			Lastmod:  h.File.ModTime.Format(time.RFC3339),
			Priority: "0.6",
		})
	}
	return entries
}

func indexHolds(s *server.Server, lang string) []*Entry {
	entries := []*Entry{}
	for _, category := range s.Trees["index"].Holds {
		holds := category.TraverseHolds()
		for _, h := range holds {
			if h.Info["translated"] == "false"  {
				continue
			}
			entries = append(entries, &Entry{
				Loc:      absoluteURL(h.Permalink(lang), lang),
				Lastmod:  h.File.ModTime.Format(time.RFC3339),
				Priority: "0.6",
			})
		}
	}
	return entries
}

func elEntries(s *server.Server, page, lang string) ([]*Entry, error) {
	entries := []*Entry{}

	els := entry.Els{}
	prio := ""

	if page == "graph" {
		els = s.Recents["graph"].ExcludePrivate()
		prio = "0.5"
	} else {
		els = s.Recents["index"]
		prio = "0.4"
	}

	for _, e := range els {
		file, err := entry.ElFileSafe(e)
		if err != nil {
			log.Println("sitemaps: els TODO")
			log.Println(err)
			continue
			return nil, err
		}
		path, err := entry.PermalinkSafe(e, lang)
		if err != nil {
			log.Println("sitemaps: els TODO")
			log.Println(err)
			continue
			return nil, err
		}
		entries = append(entries, &Entry{
			Loc:      absoluteURL(path, lang),
			Lastmod:  file.ModTime.Format(time.RFC3339),
			Priority: prio,
		})
	}
	return entries, nil
}

func absoluteURL(path, lang string) string {
	if lang == "en" {
		return fmt.Sprintf("https://en.stferal.com%v", path)
	}
	return fmt.Sprintf("https://stferal.com%v", path)
}
