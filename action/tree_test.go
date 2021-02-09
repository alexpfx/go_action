package action

import (
	"fmt"
	"github.com/alexpfx/go_action/action/builtin"
	
	"testing"
)

func TestActionTree_Show(t *testing.T) {
	
	res, err := NewFzfTree("selecione o comando", []Action{
		{
			Name:   "echo ",
			Binary: builtin.Echo,
			
			Args: []string{"alo"},
		},
		{
			Name:   "ls -la",
			Binary: builtin.Ls,
			Args:   []string{"-la"},
		},
	}).Show()
	if err != nil {
		_ = fmt.Errorf(err.Error())
	}
	fmt.Println("selected: ", res.Selected().Name)
}
