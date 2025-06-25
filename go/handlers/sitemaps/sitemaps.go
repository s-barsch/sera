package sitemaps

import (
	"fmt"
	"log"
	"net/http"

	//"g.rg-s.com/sera/go/entry/types/tree"
	"time"

	s "g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
)

type SitemapEntry struct {
	Loc      string
	Lastmod  string
	Priority string
}

func Index(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	domain := "https://seraferal.com"
	if m.Lang == "en" {
		panic("eng domain sitempas does not exist")
		//domain = "https://en.seraferal.com"
	}
	err := s.Srv.Templates.ExecuteTemplate(w, "sitemap-index", struct{ Domain string }{domain})
	if err != nil {
		log.Println(err)
		return
	}
}

func Core(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	entries, err := coreEntries(m.Lang)
	if err != nil {
		http.Error(w, "internal error", 500)
		log.Println(err)
		return
	}

	err = s.Srv.Templates.ExecuteTemplate(w, "sitemap", entries)
	if err != nil {
		log.Println(err)
		return
	}
}

func Trees(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	entries := categoryTrees(m.Lang)

	entries = append(entries, holdEntries(m.Lang)...)

	err := s.Srv.Templates.ExecuteTemplate(w, "sitemap", entries)
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

func Kines(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	entries, err := elEntries("cache", m.Lang)
	if err != nil {
		http.Error(w, "internal error", 500)
		log.Println(err)
		return
	}

	err = s.Srv.Templates.ExecuteTemplate(w, "sitemap", entries)
	if err != nil {
		log.Println(err)
		return
	}
}

func GraphEntries(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	entries, err := elEntries("graph", m.Lang)
	if err != nil {
		http.Error(w, "internal error", 500)
		log.Println(err)
		return
	}

	err = s.Srv.Templates.ExecuteTemplate(w, "sitemap", entries)
	if err != nil {
		log.Println(err)
		return
	}
}

func coreEntries(lang string) ([]*SitemapEntry, error) {
	entries := []*SitemapEntry{}

	tIndex := s.Srv.Store.Recents["indecs"].Access(false)[lang][0].Date()

	tGraph := s.Srv.Store.Recents["graph"].Access(false)[lang][0].Date()

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
			lastmod = s.Srv.Store.Trees["about"].Access(false)[lang].File().ModTime
		}

		entries = append(entries, &SitemapEntry{
			Loc:      absoluteURL(v.Href, lang),
			Priority: priority,
			Lastmod:  lastmod.Format(time.RFC3339),
		})
	}
	return entries, nil
}

func categoryTrees(lang string) []*SitemapEntry {
	entries := []*SitemapEntry{}
	trees := s.Srv.Store.Trees["graph"].Access(false)[lang].TraverseTrees()
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

func holdEntries(lang string) []*SitemapEntry {
	return aboutHolds(lang)
	//return append(indecsHolds(s, lang), aboutHolds(s, lang)...)
}

func aboutHolds(lang string) []*SitemapEntry {
	entries := []*SitemapEntry{}
	trees := s.Srv.Store.Trees["about"].Access(false)[lang].TraverseTrees()
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

func elEntries(page, lang string) ([]*SitemapEntry, error) {
	entries := []*SitemapEntry{}

	prio := ""

	//es := entry.Entries{}
	es := s.Srv.Store.Recents[page].Access(false)[lang]
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
		return fmt.Sprintf("https://en.seraferal.com%v", path)
	}
	return fmt.Sprintf("https://seraferal.com%v", path)
}
