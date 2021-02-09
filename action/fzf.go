package action

import "github.com/alexpfx/go_menus/menu"

func NewFzfTree(prompt string, actions []Action) Tree {
	
	fzfMenu := menu.NewFzfBuilder().Prompt("selecione").
		AutoSelect(true).
		Prompt(prompt).(menu.FzfBuilder).
		WithNth("2", sep).Build()
	
	return tree{
		menu:    fzfMenu,
		actions: actions,
	}
	
}
