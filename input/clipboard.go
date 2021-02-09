package input

import (
	"github.com/alexpfx/go_action/internal/util"
	"github.com/atotto/clipboard"
)

type ClipResolver struct {
}

func (c ClipResolver) Resolve(config *ResolverConfig) ([]string, error) {
	
	clipStr, _ := clipboard.ReadAll()
	splited := util.SplitSep(clipStr, config.ArgSep)
	
	res := make([]string, 0)
	
	for i, s := range splited {
		if i < len(config.Keys) {
			res = append(res, config.Keys[i])
		}
		res = append(res, s)
	}
	return res, nil
}
