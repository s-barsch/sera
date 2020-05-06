package tmpl 

import (
	"stferal/go/entry"
)

type EntryLangLazy struct {
	Entry entry.Entry
	Lang  string
	Lazy  bool
}

func NewEntryLang(e entry.Entry, lang string) *EntryLangLazy {
	return &EntryLangLazy{
		Entry: e,
		Lang:  lang,
		Lazy:  false,
	}
}

func NewEntryLangLazy(e entry.Entry, lang string, lazy bool) *EntryLangLazy {
	return &EntryLangLazy{
		Entry: e,
		Lazy:  lazy,
		Lang:  lang,
	}
}

func (e *EntryLangLazy) E() entry.Entry {
	return e.Entry
}

func (e *EntryLangLazy) L() string {
	return e.Lang
}


