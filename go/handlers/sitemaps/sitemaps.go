package sitemaps

import (
	"fmt"
	"log"
	"net/http"
	"sacer/go/entry"
	"sacer/go/head"
	"sacer/go/server"
	"sacer/go/entry/types/tree"
	"time"
)

type SitemapEntry struct {
	Loc      string
	Lastmod  string
	Priority string
}

func Index(s *server.Server, w http.ResponseWriter, r *http.Request) {
	domain := "https://sacer.site"
	if head.Lang(r.Host) == "en" {
		domain = "https://en.sacer.site"
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

func Trees(s *server.Server, w http.ResponseWriter, r *http.Request) {
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

func GraphEntries(s *server.Server, w http.ResponseWriter, r *http.Request) {
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

func coreEntries(s *server.Server, lang string) ([]*SitemapEntry, error) {
	entries := []*SitemapEntry{}

	tIndex := s.Recents["index"].Public[lang][0].Date()

	tGraph := s.Recents["graph"].Public[lang][0].Date()

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
			lastmod = s.Trees["about"].Public[lang].File().ModTime
		}

		entries = append(entries, &SitemapEntry{
			Loc:      absoluteURL(v.Href, lang),
			Priority: priority,
			Lastmod:  lastmod.Format(time.RFC3339),
		})
	}
	return entries, nil
}

func categoryEntries(s *server.Server, lang string) []*SitemapEntry {
	entries := []*SitemapEntry{}
	trees := tree.Trees{
		s.Trees["graph"].Public[lang],
		s.Trees["index"].Public[lang],
	}
	for _, tree := range trees {
		for _, t := range tree.Trees {
			if t.Level() != 1 {
				continue
			}
			entries = append(entries, &SitemapEntry{
				Loc:      absoluteURL(t.Perma(lang), lang),
				Lastmod:  t.File().ModTime.Format(time.RFC3339),
				Priority: "0.7",
			})
		}
	}
	return entries
}

func holdEntries(s *server.Server, lang string) []*SitemapEntry {
	return append(indexHolds(s, lang), aboutHolds(s, lang)...)
}

func aboutHolds(s *server.Server, lang string) []*SitemapEntry {
	entries := []*SitemapEntry{}
	trees := s.Trees["about"].Public[lang].TraverseTrees()
	for _, t := range trees {
		entries = append(entries, &SitemapEntry{
			Loc:      absoluteURL(t.Perma(lang), lang),
			Lastmod:  t.File().ModTime.Format(time.RFC3339),
			Priority: "0.6",
		})
	}
	return entries
}

func indexHolds(s *server.Server, lang string) []*SitemapEntry {
	entries := []*SitemapEntry{}
	for _, category := range s.Trees["index"].Public[lang].Trees {
		trees := category.TraverseTrees()
		for _, t := range trees {
			entries = append(entries, &SitemapEntry{
				Loc:      absoluteURL(t.Perma(lang), lang),
				Lastmod:  t.File().ModTime.Format(time.RFC3339),
				Priority: "0.6",
			})
		}
	}
	return entries
}

func elEntries(s *server.Server, page, lang string) ([]*SitemapEntry, error) {
	entries := []*SitemapEntry{}

	es := entry.Entries{}
	prio := ""

	if page == "graph" {
		es = s.Recents["graph"].Public[lang]
		prio = "0.5"
	} else {
		es = s.Recents["index"].Public[lang]
		prio = "0.4"
	}

	for _, e := range es {
		entries = append(entries, &SitemapEntry{
			Loc:      absoluteURL(e.Perma(lang), lang),
			Lastmod:  e.File().ModTime.Format(time.RFC3339),
			Priority: prio,
		})
	}
	return entries, nil
}

func absoluteURL(path, lang string) string {
	if lang == "en" {
		return fmt.Sprintf("https://en.sacer.site%v", path)
	}
	return fmt.Sprintf("https://sacer.site%v", path)
}
