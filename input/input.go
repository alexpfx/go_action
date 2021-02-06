package input

import (
	"github.com/atotto/clipboard"
	"strings"
)

type Resolver interface {
	Resolve(inputs *Config) ([]string, error)
}

type RofiResolver struct {
}

type Config struct {
	Resolver Resolver
	Keys     []string
	ArgSep   string
}

func (r RofiResolver) Resolve(config Config) ([]string, error) {
	
	res := make([]string, 0)
	
	return res, nil
}

type ClipResolver struct {
}

func (c ClipResolver) Resolve(config *Config) ([]string, error) {
	
	clipStr, _ := clipboard.ReadAll()
	splited := strings.Split(clipStr, config.ArgSep)
	
	res := make([]string, 0)
	
	for i, s := range splited {
		if i < len(config.Keys) {
			res = append(res, config.Keys[i])
		}
		res = append(res, s)
	}
	return res, nil
}
