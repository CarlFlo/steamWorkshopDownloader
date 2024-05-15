package server

import (
	"fmt"
	"regexp"
)

var (
	errNoWorkshopURLProvided = fmt.Errorf("no workshop url provided")
)

func CheckUrlInput(url string) error {

	if len(url) == 0 {
		return errNoWorkshopURLProvided
	}

	r := regexp.MustCompile(`^https:\/\/steamcommunity\.com\/sharedfiles\/filedetails\/\?id=\d+$`)
	if !r.MatchString(url) {
		return fmt.Errorf("invalid workshop URL '%s'", url)
	}

	return nil
}
