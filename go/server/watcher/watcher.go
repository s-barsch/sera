package watcher

import (
	"os"
	"os/signal"
	"strings"

	"g.rg-s.com/sacer/go/requests/tmpl"
	"github.com/rjeczalik/notify"
	"github.com/sirupsen/logrus"
)

type watcher struct {
	logger   *logrus.Logger
	trigger  chan struct{}
	notifier chan notify.EventInfo
	quitter  chan os.Signal
}

func Init(paths tmpl.Paths, trigger chan struct{}) error {
	quitter := make(chan os.Signal, 1)
	notifier := make(chan notify.EventInfo, 1)

	w := &watcher{
		trigger:  trigger,
		notifier: notifier,
		quitter:  quitter,
	}

	err := w.register(paths)
	if err != nil {
		return err
	}

	go w.watch()

	return nil
}

func (w *watcher) register(paths tmpl.Paths) error {
	signal.Notify(w.quitter, os.Interrupt)

	register := []string{
		paths.Data,
		paths.Root + "/html",
		paths.Root + "/css/dist",
		paths.Root + "/js",
	}

	for _, path := range register {
		err := notify.Watch(
			path+"...",
			w.notifier,
			notify.Remove,
			notify.Rename,
			notify.Create,
			notify.Write,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (w *watcher) watch() {
	for {
		select {
		case eventInfo := <-w.notifier:
			w.printEvent(eventInfo)
			w.trigger <- struct{}{}
		case <-w.quitter:
			notify.Stop(w.notifier)

			w.logger.Println("stopped watcher")

			os.Exit(0)
			return
		}
	}
}

func (w *watcher) printEvent(ei notify.EventInfo) {
	w.logger.Printf(
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
