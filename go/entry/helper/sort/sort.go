package sort

import (
	"stferal/go/entry"
	"io/ioutil"
	p "path/filepath"
	"strings"
	//"sort"
	"os"
)

func SortEntries(path string, entries entry.Entries) (entry.Entries, error) {
	if HasSortFile(path) {
		return ApplySortFile(path, entries)
	}

	entries.SortAsc()

	return entries, nil
}

func HasSortFile(path string) bool {
	_, err := os.Stat(SortFilePath(path))
	return err == nil
}

func SortFilePath(path string) string {
	return p.Join(path, ".sort")
}

func ApplySortFile(path string, entries entry.Entries) (entry.Entries, error) {
	sortfile, err := ioutil.ReadFile(SortFilePath(path))
	if err != nil {
		return nil, err
	}

	l := splitToLines(string(sortfile))

	for _, sortElement := range invert(l) {
		for i, e := range entries {
			if e.File().Name() == sortElement {
				cut := entries[i]
				entries = append(entry.Entries{cut}, append(entries[:i], entries[i+1:]...)...)
			}
		}
	}
	return entries, nil
}


func splitToLines(file string) []string {
	return strings.Split(strings.TrimSpace(file), "\n")
}

func invert(ss []string) []string {
	ns := []string{}
	for i := len(ss) - 1; i >= 0; i-- {
		ns = append(ns, ss[i])
	}
	return ns
}


