package index

/*
import (
	"log"
	"net/http"
	"stferal/go/entry"
	"stferal/go/head"
	"stferal/go/server"
	"stferal/go/entry/types/tree"
)

type indexPage struct {
	Head *head.Head
	Entry entry.Entry
}

func IndexPage(s *server.Server, w http.ResponseWriter, r *http.Request, tr *tree.Tree) {
	if perma := tr.Perma(head.Lang(r.Host)); r.URL.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}
	head := &head.Head{
		Title:   indexTitle(tr, head.Lang(r.Host)),
		Section: "index",
		Path:    r.URL.Path,
		Host:    r.Host,
		Entry:   tr,
		Options: head.GetOptions(r),
	}
	err := head.Make()
	if err != nil {
		s.Log.Println(err)
		return
	}

	err = s.ExecuteTemplate(w, "index-main", &indexPage{
		Head: head,
		Entry: tr,
	})
	if err != nil {
		log.Println(err)
	}
}

func indexTitle(tr *tree.Tree, lang string) string {
	title := tr.Info().Title(lang)

	c := tr.Chain(lang)
	if len(c) > 2 {
		// TODO:
		//title += " - " + c[1].Title
		title += " - " + c[1]
	}

	return title
}
*/
