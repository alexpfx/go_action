package input

import "io"

type Resolver interface {
	Resolve(inputs *ResolverConfig) ([]string, error)
}

type ResolverConfig struct {
	Resolver Resolver
	Keys     []string
	ArgSep   string
	Reader io.Reader
}
