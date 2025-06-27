package cache

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"g.rg-s.com/sera/go/entry"
	"g.rg-s.com/sera/go/entry/tools"
	s "g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
	"g.rg-s.com/sera/go/viewer"
)

type cacheSingle struct {
	Meta      *meta.Meta
	Entry     entry.Entry
	Neighbors []entry.Entry
}

func ServeSingle(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cache := s.Srv.Store.Trees["cache"].Access(m.Auth.Subscriber)[m.Lang]
		e, err := cache.LookupEntryHash(m.Split.Hash)
		if err != nil {
			http.Redirect(w, r, "/cache", http.StatusMovedPermanently)
			return
		}

		perma := e.Perma(m.Lang)
		if m.Path != perma {
			http.Redirect(w, r, perma, http.StatusMovedPermanently)
			return
		}

		m.Title = getTitle(e, m.Lang)
		m.SetSection("cache")
		m.SetHreflang(e)

		err = s.Srv.ExecuteTemplate(w, "cache-single", &cacheSingle{
			Meta:  m,
			Entry: e,
		})
		if err != nil {
			log.Println(err)
		}
	}
}

func getDate(d time.Time, lang string) string {
	return fmt.Sprintf(d.Format("02 %v 2006"), tools.Abbr(tools.MonthLang(d, lang)))
}

func getTitle(e entry.Entry, lang string) string {
	return fmt.Sprintf("%v - %v - %v", e.Title(lang), getDate(e.Date(), lang), "Cache")
}

/*
func getNeighbors(es entry.Entries, hash string) []entry.Entry {
	cpy := make([]entry.Entry, len(es))
	copy(cpy, es)
	for i, e := range cpy {
		if e.Hash() == hash {
			l := len(es)
			j, k := limits(l, i)
			d := i + 1
			if d > l {
				d = i
			}
			return append(cpy[j:i], cpy[d:k]...)
		}
	}
	return nil
}
*/

/*
// TODO: verbose... should be simplified
func limits(l, i int) (int, int) {
	// number of neighors left and right
	number := 2
	j, k := 0, l
	// set start position
	if x := i - number; x > 0 {
		j = x
	}

	if left := i - j; left < 2 {
		number = number + (number - left)
	}

	if y := i + 1 + number; y < l {
		k = y
	} else {
		k = l
	}

	if right := k - i - 1; right < 2 {
		if j-(number-right) > 0 {
			j = j - (number - right)
		} else {
			j = 0
		}
	}

	return j, k
}
*/
