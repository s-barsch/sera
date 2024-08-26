package reels

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"g.sacerb.com/sacer/go/entry"
	"g.sacerb.com/sacer/go/entry/tools"
	"g.sacerb.com/sacer/go/server"
	"g.sacerb.com/sacer/go/server/meta"
	"g.sacerb.com/sacer/go/server/paths"
)

type reelsSingle struct {
	Meta      *meta.Meta
	Entry     entry.Entry
	Neighbors []entry.Entry
}

func ServeSingle(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta, p *paths.Path) {
	reels := s.Trees["reels"].Access(m.Auth.Subscriber)[m.Lang]
	e, err := reels.LookupEntryHash(p.Hash)
	if err != nil {
		http.Redirect(w, r, "/reels", 301)
		return
	}

	perma := e.Perma(m.Lang)
	if m.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}

	m.Title = getTitle(e, m.Lang)
	m.Section = "reels"

	err = m.Process(e)
	if err != nil {
		s.Log.Println(err)
		return
	}

	err = s.ExecuteTemplate(w, "reels-single", &reelsSingle{
		Meta:  m,
		Entry: e,
		//Neighbors: getNeighbors(s.Recents["reels"].Access(m.Auth.Subscriber)[m.Lang], p.Hash),
	})
	if err != nil {
		log.Println(err)
	}
}

func getDate(d time.Time, lang string) string {
	return fmt.Sprintf(d.Format("02 %v 2006"), tools.Abbr(tools.MonthLang(d, lang)))
}

func getTitle(e entry.Entry, lang string) string {
	return fmt.Sprintf("%v - %v - %v", e.Title(lang), getDate(e.Date(), lang), strings.Title(tools.KineName[lang]))
}

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
