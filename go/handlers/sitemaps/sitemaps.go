package sitemaps

import (
	"fmt"
	"log"
	"net/http"

	//"g.rg-s.com/sacer/go/entry/types/tree"
	"time"

	"g.rg-s.com/sacer/go/server"
	"g.rg-s.com/sacer/go/server/meta"
)

type SitemapEntry struct {
	Loc      string
	Lastmod  string
	Priority string
}

func Index(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	domain := "https://sacer.site"
	if m.Lang == "en" {
		domain = "https://en.sacer.site"
	}
	err := s.Templates.ExecuteTemplate(w, "sitemap-index", struct{ Domain string }{domain})
	if err != nil {
		log.Println(err)
		return
	}
}

func Core(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	entries, err := coreEntries(s, m.Lang)
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

func Trees(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	entries := categoryTrees(s, m.Lang)

	entries = append(entries, holdEntries(s, m.Lang)...)

	err := s.Templates.ExecuteTemplate(w, "sitemap", entries)
	if err != nil {
		log.Println(err)
		return
	}
}

/*
func IndexEls(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	entries, err := smEls("indecs", lang(r.Host))
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

func Kines(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	entries, err := elEntries(s, "cache", m.Lang)
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

func GraphEntries(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	entries, err := elEntries(s, "graph", m.Lang)
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

	tIndex := s.Recents["indecs"].Access(false)[lang][0].Date()

	tGraph := s.Recents["graph"].Access(false)[lang][0].Date()

	for _, v := range meta.NewNav(lang) {
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
		case "indecs":
			lastmod = tIndex
		case "graph":
			lastmod = tGraph
		case "about":
			lastmod = s.Trees["about"].Access(false)[lang].File().ModTime
		}

		entries = append(entries, &SitemapEntry{
			Loc:      absoluteURL(v.Href, lang),
			Priority: priority,
			Lastmod:  lastmod.Format(time.RFC3339),
		})
	}
	return entries, nil
}

func categoryTrees(s *server.Server, lang string) []*SitemapEntry {
	entries := []*SitemapEntry{}
	trees := s.Trees["graph"].Access(false)[lang].TraverseTrees()
	/*
		trees := tree.Trees{
			//s.Trees["indecs"].Access(false)[lang],
		}
	*/
	for _, t := range trees {
		entries = append(entries, &SitemapEntry{
			Loc:      absoluteURL(t.Perma(lang), lang),
			Lastmod:  t.File().ModTime.Format(time.RFC3339),
			Priority: "0.7",
		})
	}
	return entries
}

func holdEntries(s *server.Server, lang string) []*SitemapEntry {
	return aboutHolds(s, lang)
	//return append(indecsHolds(s, lang), aboutHolds(s, lang)...)
}

func aboutHolds(s *server.Server, lang string) []*SitemapEntry {
	entries := []*SitemapEntry{}
	trees := s.Trees["about"].Access(false)[lang].TraverseTrees()
	for _, t := range trees {
		entries = append(entries, &SitemapEntry{
			Loc:      absoluteURL(t.Perma(lang), lang),
			Lastmod:  t.File().ModTime.Format(time.RFC3339),
			Priority: "0.6",
		})
	}
	return entries
}

/*
func indecsHolds(s *server.Server, lang string) []*SitemapEntry {
	entries := []*SitemapEntry{}
	for _, category := range s.Trees["indecs"].Access(false)[lang].Trees {
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
*/

func elEntries(s *server.Server, page, lang string) ([]*SitemapEntry, error) {
	entries := []*SitemapEntry{}

	prio := ""

	//es := entry.Entries{}
	es := s.Recents[page].Access(false)[lang]
	prio = "0.5"

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
