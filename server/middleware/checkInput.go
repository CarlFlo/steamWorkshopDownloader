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
	} else if len(url) > 256 {
		// A "normal" url is around 65 characters
		return errProvidedURLTooLong
	}

	// the url is later processed using the net/url package to extract the id from the url

	r := regexp.MustCompile(`^https:\/\/steamcommunity\.com\/sharedfiles\/filedetails\/\?id=\d+`)
	if !r.MatchString(url) {
		return fmt.Errorf("invalid workshop URL '%s'", url)
	}

	return nil
}
