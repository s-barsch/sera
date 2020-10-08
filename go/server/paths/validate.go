package paths

import (
	"fmt"
	"regexp"
	"strconv"
)

// TODO: Aren’t there better functions for this?

var validPath = regexp.MustCompile(`^\/[0-9a-z+-_.\/]*$`)
var validName = regexp.MustCompile(`^[0-9a-z+-_.]*$`)

func Sanitize(foreign string) (string, error) {
	if validPath.MatchString(foreign) {
		return foreign, nil
	}
	return "", fmt.Errorf("invalid Path: %v", foreign)
}

func SanitizeName(foreign string) (string, error) {
	if validName.MatchString(foreign) {
		return foreign, nil
	}
	return "", fmt.Errorf("invalid Path: %v", foreign)
}

func SanitizeInt(foreign string) (int, error) {
	return strconv.Atoi(foreign)
}
