package entry

import (
	"fmt"
	"time"
)

func Title(e interface{}, lang string) string {
	switch e.(type) {
	case *Text:
		return e.(*Text).Title(lang)
	case *Image:
		return e.(*Image).Title(lang)
	case *Video:
		return e.(*Video).Title(lang)
	case *Audio:
		return e.(*Audio).Title(lang)
	case *Html:
		return e.(*Html).Title(lang)
	case *Set:
		return e.(*Set).Title(lang)
	default:
		return ""
	}
}

func ElTitleSafe(e interface{}, lang string) (string, error) {
	return "", nil
}

func Id(e interface{}) string {
	switch e.(type) {
	case *Text:
		return e.(*Text).Date.Format(Timestamp)
	case *Image:
		return e.(*Image).Date.Format(Timestamp)
	case *Video:
		return e.(*Video).Date.Format(Timestamp)
	case *Audio:
		return e.(*Audio).Date.Format(Timestamp)
	case *Set:
		return e.(*Set).Date.Format(Timestamp)
	case *Hold:
		return e.(*Hold).Date.Format(Timestamp)
	case *Html:
		return e.(*Html).Date.Format(Timestamp)
	case *File:
		return e.(*File).Id
	default:
		return ""
	}
}

func ElHold(e interface{}) (*Hold, error) {
	switch e.(type) {
	case *Text:
		return e.(*Text).File.Hold, nil
	case *Image:
		return e.(*Image).File.Hold, nil
	case *Video:
		return e.(*Video).File.Hold, nil
	case *Audio:
		return e.(*Audio).File.Hold, nil
	case *Set:
		return e.(*Set).File.Hold, nil
	case *Html:
		return e.(*Html).File.Hold, nil
	case *File:
		return e.(*File).Hold, nil
	case *Hold:
		return e.(*Hold).Mother, nil
	default:
		return nil, fmt.Errorf("Hold not found. %v", e)
	}
}

func ElFile(e interface{}) *File {
	switch e.(type) {
	case *Text:
		return e.(*Text).File
	case *Image:
		return e.(*Image).File
	case *Video:
		return e.(*Video).File
	case *Audio:
		return e.(*Audio).File
	case *Set:
		return e.(*Set).File
	case *File:
		return e.(*File)
	case *Hold:
		return e.(*Hold).File
	case *Html:
		return e.(*Html).File
	default:
		return nil
	}
}

func ElFileSafe(e interface{}) (*File, error) {
	f := ElFile(e)
	if f == nil {
		return nil, fmt.Errorf("File not found. %v", e)
	}
	return f, nil
}

func ElSection(eh interface{}) string {
	h, ok := eh.(*Hold)
	if ok {
		return h.Section()
	} else {
		f, err := ElFileSafe(eh)
		if err != nil {
			// err?
			return ""
		}
		return f.Section()
	}
	return ""
}

func ElSectionSafe(eh interface{}) (string, error) {
	page := ElSection(eh)
	if page == "" {
		return "", fmt.Errorf("Section not found for El/Hold: %v", eh)
	}
	return page, nil
}

func Type(e interface{}) string {
	switch e.(type) {
	case *Text:
		return "text"
	case *Image:
		return "image"
	case *Set:
		return "set"
	case *File:
		return "file"
	case *Audio:
		return "audio"
	case *Video:
		return "video"
	case *Hold:
		return "hold"
	case *Html:
		return "html"
	default:
		return ""
	}
}

func Permalink(e interface{}, lang string) string {
	switch e.(type) {
	case *Image:
		return e.(*Image).Permalink(lang)
	case *Set:
		return e.(*Set).Permalink(lang)
	case *Text:
		return e.(*Text).Permalink(lang)
	case *Audio:
		return e.(*Audio).Permalink(lang)
	case *Video:
		return e.(*Video).Permalink(lang)
	case *Hold:
		return e.(*Hold).Permalink(lang)
	case *Html:
		return e.(*Html).Permalink(lang)
	default:
		return ""
	}
}

func PermalinkSafe(e interface{}, lang string) (string, error) {
	perma := Permalink(e, lang)
	if perma == "" {
		return "", fmt.Errorf("Element has no permalink. %v", e)
	}
	return perma, nil
}

func Date(e interface{}) time.Time {
	switch e.(type) {
	case *Image:
		return e.(*Image).Date
	case *Set:
		return e.(*Set).Date
	case *Text:
		return e.(*Text).Date
	case *Audio:
		return e.(*Audio).Date
	case *Video:
		return e.(*Video).Date
	case *Html:
		return e.(*Html).Date
	default:
		return time.Time{}
	}
}

func DateSafe(e interface{}) (time.Time, error) {
	d := Date(e)
	if d.IsZero() {
		return d, fmt.Errorf("Date of element is Zero. %v", e)
	}
	return d, nil
}

func ModTime(e interface{}) time.Time {
	switch e.(type) {
	case *Hold:
		return e.(*Hold).File.ModTime
	case *Image:
		return e.(*Image).File.ModTime
	case *Set:
		return e.(*Set).File.ModTime
	case *Text:
		return e.(*Text).File.ModTime
	case *File:
		return e.(*File).ModTime
	case *Audio:
		return e.(*Audio).File.ModTime
	case *Video:
		return e.(*Video).File.ModTime
	case *Html:
		return e.(*Html).File.ModTime
	default:
		return time.Time{}
	}
}

func ModTimeSafe(e interface{}) (time.Time, error) {
	t := ModTime(e)
	if t.IsZero() {
		return t, fmt.Errorf("ModTime of Element is Zero. %v", e)
	}
	return t, nil
}

func InfoSafe(eh interface{}) Info {
	switch eh.(type) {
	case *Hold:
		return eh.(*Hold).Info
	case *Image:
		return eh.(*Image).Info
	case *Set:
		return eh.(*Set).Info
	case *Text:
		return eh.(*Text).Info
	case *Audio:
		return eh.(*Audio).Info
	case *Video:
		return eh.(*Video).Info
	case *Html:
		return eh.(*Html).Info
	default:
		return map[string]string{}
	}
}

func EntryInfo(eh interface{}) (Info, error) {
	switch eh.(type) {
	case *Hold:
		return eh.(*Hold).Info, nil
	case *Image:
		return eh.(*Image).Info, nil
	case *Set:
		return eh.(*Set).Info, nil
	case *Text:
		return eh.(*Text).Info, nil
	case *Audio:
		return eh.(*Audio).Info, nil
	case *Video:
		return eh.(*Video).Info, nil
	case *Html:
		return eh.(*Html).Info, nil
	default:
		return nil, fmt.Errorf("Info: type not found. %v", eh)
	}
}

func setInfo(eh interface{}, i Info) error {
	switch eh.(type) {
	case *Hold:
		eh.(*Hold).Info = i
	case *Image:
		eh.(*Image).Info = i
	case *Set:
		eh.(*Set).Info = i
	case *Text:
		eh.(*Text).Info = i
	case *Audio:
		eh.(*Audio).Info = i
	case *Video:
		eh.(*Video).Info = i
	case *Html:
		eh.(*Html).Info = i
	default:
		return fmt.Errorf("Info: type not found. %v", eh)
	}
	return nil
}
