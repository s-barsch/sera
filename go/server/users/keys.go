package users

import (
	"encoding/base64"
	"strings"
	"fmt"
)

func EncodeMailKey(mail, key string) string {
	return base64.StdEncoding.EncodeToString([]byte(mail + "+" + key))
}

func DecodeMailKey(b64 string) (mail, key string, err error) {
	b, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return "", "", err
	}
	str := string(b)
	i := strings.Index(str, "+")
	if i <= 0 {
		err = fmt.Errorf("decodeVerify: cannot get key and string")
		return "", "", err
	}
	return str[:i], str[i+1:], nil

}

func GenerateSessionKey() (string, error) {
	return GenerateRandomStringURLSafe(40)
}

func GenerateLoginKey() (string, error) {
	return GenerateRandomStringURLSafe(32)
}

