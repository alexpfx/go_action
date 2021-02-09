package action

import (
	"fmt"
	"github.com/alexpfx/go_action/action/binary"
	
	"testing"
)

func TestActionTree_Show(t *testing.T) {
	
	res, err := NewFzfTree("selecione o comando", []Action{
		{
			Name:   "echo ",
			Binary: binary.Echo,
			
			Args: []string{"alo"},
		},
		{
			Name:   "ls -la",
			Binary: binary.Ls,
			Args:   []string{"-la"},
		},
	}).Show()
	if err != nil {
		_ = fmt.Errorf(err.Error())
	}
	fmt.Println("selected: ", res.Selected().Name)
}
