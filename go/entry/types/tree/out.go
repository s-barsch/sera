package tree

import (
	"fmt"
	"sacer/go/entry"
	"sacer/go/entry/tools"
)

func (t *Tree) CombinedTitle(lang string) string {
	topicPage := ""
	if topic := t.TopicPage(); topic != nil {
		topicPage = topic.Info().Title(lang) + " – "
	}
	return topicPage + t.Title(lang)
}

func (t *Tree) Label(lang string) string {
	if label := t.info.Field("label", lang); label != "" {
		return label
	}
	return t.Title(lang)
}

// For this to work, title fields have to be set for all sections. Otherwise,
// the root node will return a short hash.
func (t *Tree) Section() string {
	section := t.Chain()[0].Slug("en")
	if section == "cine" {
		return "kine"
	}
	return section
}

func (t *Tree) Perma(lang string) string {
	switch t.Section() {
	case "graph":
		return graphPerma(t, lang)
	case "kine":
		return kinePerma(t, lang)
	case "indecs":
		return indecsPerma(t, lang)
	case "about":
		return aboutPerma(t, lang)
	}
	return fmt.Sprintf("/%v", t.Slug(lang))
}

func aboutPerma(t *Tree, lang string) string {
	if t.Level() < 2 {
		return t.Path(lang)
	}
	return defaultPerma(t, lang)
}

func defaultPerma(t *Tree, lang string) string {
	return fmt.Sprintf("%v/%v-%v", t.parent.Path(lang), t.Title(lang), t.Hash())
}

func extraPerma(t *Tree, lang string) string {
	return fmt.Sprintf("/%v", t.Title(lang))
}

func kinePerma(t *Tree, lang string) string {
	start := fmt.Sprintf("/%v/%v", lang, tools.KineName[lang])
	switch l := t.Level(); {
	case l == 0:
		return start
		/*
			case l == 1:
				last := start
				year := t
				if d := len(year.Trees); d > 0 {
					month := year.Trees[d-1]
					if x := len(month.Entries()); x > 0 {
						last = month.Entries()[x-1].Perma(lang)
					}
				}
				return last
		*/
	case l == 2:
		last := start
		if l := len(t.Entries()); l > 0 {
			last = t.Entries()[l-1].Perma(lang)
		}
		return last
	case l < 3:
		return t.Path(lang)
	}
	return fmt.Sprintf("/graph-permalink-error-%v", t.Title(lang))
}

func graphPerma(t *Tree, lang string) string {
	switch l := t.Level(); {
	case l == 0:
		return fmt.Sprintf("%v/graph", lang)
	case l == 2:
		fallthrough
	case l < 3:
		return t.Path(lang)
	}
	return fmt.Sprintf("/graph-permalink-error-%v", t.Title(lang))
}

func indecsPerma(t *Tree, lang string) string {
	switch l := t.Level(); {
	case l == 0:
		return "/indecs"
	case l < 2:
		// prints /indecs/welt
		return t.Path(lang)
	}
	// prints /indecs/welt/wuestenleben-36c35dcb
	return fmt.Sprintf("%v-%v", t.Path(lang), t.Hash())
}

func (t *Tree) Path(lang string) string {
	path := "/" + lang
	for _, tree := range t.Chain() {
		path += "/" + tree.Slug(lang)
	}
	return path
}

func (t *Tree) Chain() Trees {
	parent := typeCheck(t.Parent())
	if parent == nil {
		return Trees{t}
	}

	return append(parent.Chain(), t)
}

func typeCheck(parentEntry entry.Entry) *Tree {
	parent, ok := parentEntry.(*Tree)
	if !ok {
		return nil
	}

	return parent
}

func (t *Tree) Level() int {
	parent := typeCheck(t.Parent())
	if parent == nil {
		return 0
	}
	return 1 + parent.Level()
}

/*
func (t *Tree) IsSubPage() bool {
	parent := typeCheck(t.Parent())
	if parent == nil {
		return false
	}
	if t.parent.Info()["istopic"] == "true" {
		return true
	}
	return parent.IsSubPage()
}
*/

func (t *Tree) TopicPage() *Tree {
	parent := typeCheck(t.Parent())
	if parent == nil {
		return nil
	}
	if t.parent.Info()["istopic"] == "true" {
		return parent
	}
	return parent.TopicPage()
}
