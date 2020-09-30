package shortener

import (
	"errors"
	"net/url"
)

func validateRequest(req *Request) error {
	if req.URL == "" {
		return errors.New("url is required")
	}

	if !validURL(req.URL) {
		return errors.New("invalid url")
	}

	return nil
}

func validURL(s string) bool {
	_, err := url.ParseRequestURI(s)
	if err != nil {
		return false
	}

	u, err := url.Parse(s)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
