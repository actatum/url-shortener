package shortener

import (
	"errors"
	"regexp"
	"strings"

	"github.com/teris-io/shortid"
	"go.uber.org/zap"
)

var (
	slugRegex = regexp.MustCompile(`(?i)[\w\-]`)
)

type service struct {
	repo   *repository
	logger *zap.Logger
}

func newService(r *repository, l *zap.Logger) *service {
	return &service{
		repo:   r,
		logger: l,
	}
}

// Create a shortened url and stores it in the repository
func (s *service) Create(req *Request) (*Response, error) {
	err := validateRequest(req)
	if err != nil {
		return nil, err
	}

	if req.Slug == "" {
		req.Slug, err = shortid.Generate()
		if err != nil {
			return nil, err
		}
	}
	req.Slug = strings.ToLower(req.Slug)

	match, err := regexp.MatchString(slugRegex.String(), req.Slug)
	if err != nil {
		return nil, err
	}

	if !match {
		return nil, errors.New("invalid slug")
	}

	resp, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	s.logger.Info("url: " + req.URL + " saved at slug: " + req.Slug)
	return resp, nil
}

// Read a shortened url given a full url
func (s *service) Read(req *Request) (*Response, error) {
	resp, err := s.repo.Read(req)
	if err != nil {
		return nil, err
	}

	s.logger.Info("retrieved url: " + resp.URL + " from slug: " + req.Slug)
	return resp, nil
}

// Top5 gets the top 5 shortened urls by usage
func (s *service) Top5() ([]Response, error) {
	resp, err := s.repo.Top5()
	if err != nil {
		return nil, err
	}

	s.logger.Info("retrieved top 5 slugs")
	return resp, nil
}
