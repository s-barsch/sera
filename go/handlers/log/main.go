package log

import (
	"log"
	"net/http"
	"sacer/go/entry"
	"sacer/go/server"
	"sacer/go/server/auth"
	"sacer/go/server/head"
)

type logMain struct {
	Head     *head.Head
	Log      entry.Entries
}

func Main(s *server.Server, w http.ResponseWriter, r *http.Request, a *auth.Auth) {
	head := &head.Head{
		Title:   "",
		Section: "extra",
		Path:    "/",
		Host:    r.Host,
		Entry:   nil,
		Options: head.GetOptions(r),
	}
	err := head.Process()
	if err != nil {
		return
	}

	err = s.ExecuteTemplate(w, "log-main", &logMain{
		Head:  head,
		Log:   s.Recents["log"].Access(true)["de"],
	})
	if err != nil {
		log.Println(err)
	}
}


