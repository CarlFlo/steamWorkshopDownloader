package middleware

import (
	"fmt"
	"regexp"
)

var (
	errNoWorkshopURLProvided = fmt.Errorf("no workshop url provided")
	errProvidedURLTooLong    = fmt.Errorf("provided workshop url was too long")
)

func CheckUrlInput(url string) error {

	if len(url) == 0 {
		return errNoWorkshopURLProvided
	} else if len(url) > 512 {
		// A "normal" url is around 65 characters
		return errProvidedURLTooLong
	}

	// the url is later processed using the net/url package to extract the id from the url

	patterns := []string{
		`^https:\/\/steamcommunity\.com\/sharedfiles\/filedetails\/\?id=\d+`,
		`^(https?:\/\/)?steamcommunity\.com\/sharedfiles\/filedetails\/\?id=\d+`, // everything before 'steamcommunity' is optinal
	}

	r := regexp.MustCompile(patterns[0])
	if !r.MatchString(url) {
		return fmt.Errorf("invalid workshop URL '%s'", url)
	}

	return nil
}
