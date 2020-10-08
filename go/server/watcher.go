package server

import (
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/rjeczalik/notify"
)

func (s *Server) SetupWatcher() error {
	s.Quit = make(chan os.Signal, 1)
	s.Watcher = make(chan notify.EventInfo, 1)

	signal.Notify(s.Quit, os.Interrupt)

	paths := []string{
		s.Paths.Data,
		s.Paths.Root + "/html",
		s.Paths.Root + "/css/dist",
		s.Paths.Root + "/js/dist",
	}

	for _, path := range paths {
		err := notify.Watch(
			path+"...",
			s.Watcher,
			notify.Remove,
			notify.Rename,
			notify.Create,
			notify.Write,
		)
		if err != nil {
			return err
		}
	}

	if s.Flags.Debug {
		log.Println("Started watcher.")
	}

	go s.Watch()

	return nil
}

func runLoad(s *Server) {
	select {
	case s.Queue <- 1:
		log.Println("Load started.")
	default:
		log.Println("Blocked.")
		return
	}
	go func() {
		err := s.Load()
		if err != nil {
			log.Println(err)
		}
		<-s.Queue
	}()
}

func (s *Server) Watch() {
	for {
		select {
		case ei := <-s.Watcher:
			printEvent(ei)
			runLoad(s)
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

func printEvent(ei notify.EventInfo) {
	log.Printf(
		"%v: %v\n",
		formatEvent(ei.Event().String()),
		formatPath(ei.Path()),
	)
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
