package ezoci

import "errors"

type ComparisonResult struct {
	Child            bool
	Parents          []string
	CommonLayers     []string
	WeakCommonLayers []string
}

func Compare(source string, others ...string) (*ComparisonResult, error) {
	return &ComparisonResult{}, errors.ErrUnsupported
}
