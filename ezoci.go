package ezoci

import (
	"github.com/calamity-m/ezoci/pkg/search"
)

func Search(opts *search.SearchOpts) (string, error) {

	resolver := search.SearchResolver{Opts: opts}

	return resolver.ResolveSearch()
}
