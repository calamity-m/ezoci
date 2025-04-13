package ezoci

import (
	"errors"
	"net/http"

	"github.com/calamity-m/ezoci/pkg/providers"
)

type SearchOpts struct {
	Client      *http.Client
	Provider    providers.Provider
	UrlOverride string
	Path        string
}

func Search(opts *SearchOpts) (string, error) {
	return "", errors.ErrUnsupported
}
