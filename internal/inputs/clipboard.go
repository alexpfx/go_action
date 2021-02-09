package inputs

import (
	"github.com/alexpfx/go_action/input"
	"github.com/atotto/clipboard"
	"strings"
)

type ClipResolver struct {
}

func (c ClipResolver) Resolve(config *input.ResolverConfig) ([]string, error) {
	
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
