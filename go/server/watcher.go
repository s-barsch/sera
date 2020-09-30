package server

import (
	"github.com/rjeczalik/notify"
	"log"
	"os"
	"os/signal"
	"strings"
)

func (s *Server) SetupWatcher() error {
	s.Quit = make(chan os.Signal, 1)
	s.Watcher = make(chan notify.EventInfo, 1)

	signal.Notify(s.Quit, os.Interrupt)

	err := notify.Watch(
		s.Paths.Data+"...",
		s.Watcher,
		notify.Remove,
		notify.Rename,
		notify.Create,
		notify.Write,
	)

	if err != nil {
		return err
	}

	if s.Flags.Debug {
		log.Println("Started watcher.")
	}

	go s.Watch()

	return nil
}

func (s *Server) Watch() {
	for {
		select {
		case ei := <-s.Watcher:
			log.Printf("%v: %v", formatEvent(ei.Event().String()), formatPath(ei.Path()))
			s.Reload()
		case <-s.Quit:
			notify.Stop(s.Watcher)

			if s.Flags.Debug {
				log.Println("Stopped watcher.")
			}

			os.Exit(0)
			return
		}
	}
}

func formatPath(path string) string {
	s := strings.Split(path, "/")
	if l := len(s); l > 2 {
		return strings.Join(s[l-2:], "/")
	}
	return path
}

const notifyLen = len("notify.")

func formatEvent(event string) string {
	if len(event) > notifyLen {
		return event[notifyLen:]
	}
	return event
}
