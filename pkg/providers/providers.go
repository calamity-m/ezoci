package providers

type Provider int

const (
	Docker Provider = iota
	Harbor
	Ghcr
	Gitlab
	Gcr
)
