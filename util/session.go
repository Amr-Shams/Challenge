package util

import (
	"fmt"
	"log"
	"os"
	"github.com/browserutils/kooky/browser/safari"
)

func GetSession() (string, error) {
	return GetSessionForDomain(".adventofcode.com")
}

func GetSessionForDomain(domain string) (string, error) {
	dir, _ := os.UserHomeDir()
	cookiesFile := dir + "/Library/Containers/com.apple.Safari/Data/Library/Cookies/Cookies.binarycookies"
	cookies, err := safari.ReadCookies(cookiesFile)
	if err != nil {
		log.Fatal(err)
	}
	for _, cookie := range cookies {
		if cookie.Domain == domain && cookie.Name == "session" {
			return cookie.Value, nil
		}
	}
	return "", fmt.Errorf("session cookie not found")
}