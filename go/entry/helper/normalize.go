package helper

import (
	"github.com/kennygrant/sanitize"
	"regexp"
	"strings"
)

var validSlug = regexp.MustCompile("[^a-z0-9-]+")

func Normalize(name string) string {
	name = strings.TrimSpace(name)
	name = strings.Replace(name, "⹀", "-", -1)
	name = strings.Replace(name, " ", "-", -1)
	name = strings.ToLower(name)
	name = sanitize.Accents(name)

	name = validSlug.ReplaceAllString(name, "")
	return name
}
