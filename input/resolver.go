package input

import (
	"io"
	"regexp"
)

type Resolver interface {
	Resolve(inputs *ResolverConfig) ([]string, error)
}

type ResolverConfig struct {
	Resolver Resolver
	Keys     []string
	ArgSep   string
	Reader   io.Reader
	Prompt   string
	searchPattern regexp.Regexp
	
}
