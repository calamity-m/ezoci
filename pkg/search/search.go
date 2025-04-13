package search

import (
	"errors"
	"net/http"
)

type Provider int

const (
	Docker Provider = iota
	Harbor
	Ghcr
	Gitlab
	Gcr
)

type SearchOpts struct {
	Client      *http.Client
	Provider    Provider
	UrlOverride string
	Path        string
}

type SearchResolver struct {
	Opts *SearchOpts
}

func (s *SearchResolver) ResolveSearch() (string, error) {
	return "", errors.ErrUnsupported
}
