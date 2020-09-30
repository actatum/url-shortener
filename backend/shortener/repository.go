package shortener

import (
	"errors"
	"log"
)

type repository struct {
	cache map[string]string
}

func newRepository() *repository {
	return &repository{
		cache: make(map[string]string),
	}
}

// Create stores a key value pair of slug and url
func (r *repository) Create(req *Request) (*Response, error) {
	url, ok := r.cache[req.Slug]
	if ok {
		log.Println(url, req.Slug)
		return nil, errors.New("url with slug: " + req.Slug + " already exists")
	}
	r.cache[req.Slug] = req.URL

	return &Response{Slug: req.Slug, URL: req.URL}, nil
}

// Read reads a shortened url given a key of full url
func (r *repository) Read(req *Request) (*Response, error) {
	url, ok := r.cache[req.Slug]
	if !ok {
		return nil, errors.New("no url for slug: " + req.Slug)
	}

	return &Response{
		Slug: req.Slug,
		URL:  url,
	}, nil
}
