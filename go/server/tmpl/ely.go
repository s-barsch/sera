package tmpl

import (
	"sacer/go/entry"
	"sacer/go/entry/types/text"
)

type NotesLang struct {
	Notes *text.Notes
	Lang  string
}

func (nL NotesLang) N() *text.Notes {
	return nL.Notes
}

func (nL NotesLang) L() string {
	return nL.Lang
}

func NewNotesLang(n *text.Notes, lang string) *NotesLang {
	return &NotesLang{
		Notes: n,
		Lang:  lang,
	}
}

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
		Lang:  lang,
		Lazy:  lazy,
	}
}

func (e *EntryLangLazy) E() entry.Entry {
	return e.Entry
}

func (e *EntryLangLazy) L() string {
	return e.Lang
}

func (e *EntryLangLazy) Y() bool {
	return e.Lazy
}

type EntriesLangLazy struct {
	Entries entry.Entries
	Lang    string
	Lazy    bool
}

func NewEntriesLang(es entry.Entries, lang string) *EntriesLangLazy {
	return &EntriesLangLazy{
		Entries: es,
		Lang:    lang,
		Lazy:    false,
	}
}

func NewEntriesLangLazy(es entry.Entries, lang string, lazy bool) *EntriesLangLazy {
	return &EntriesLangLazy{
		Entries: es,
		Lazy:    lazy,
		Lang:    lang,
	}
}

func (e *EntriesLangLazy) Es() entry.Entries {
	return e.Entries
}

func (e *EntriesLangLazy) L() string {
	return e.Lang
}

func (e *EntriesLangLazy) Y() bool {
	return e.Lazy
}
