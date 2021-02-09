package input

import (
	"github.com/alexpfx/go_action/internal/util"
	"github.com/alexpfx/go_menus/menu"
)

type RofiResolver struct {
}

func (r RofiResolver) Resolve(config *ResolverConfig) ([]string, error) {
	rofiMenu := menu.NewRofiInputBuilder(config.Prompt).Build()
	
	rofiStr, err := rofiMenu.Run("")
	if err != nil {
		return nil, err
	}
	
	splited := util.SplitSep(rofiStr, config.ArgSep)
	
	res := make([]string, 0)
	
	for i, s := range splited {
		if i < len(config.Keys) {
			res = append(res, config.Keys[i])
		}
		res = append(res, s)
	}
	return res, nil
	
}
