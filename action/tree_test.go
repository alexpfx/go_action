package action

import (
	"fmt"
	"testing"
)

func TestActionTree_Show(t *testing.T) {
	ls := Binary{
		CmdPath: "ls",
		Name:    "",
		FixArgs: []string{},
	}

	res, err := NewFzfTree("selecione o comando", []Action{
		{
			Name:   "ls -la",
			Binary: ls,
			
			Args:   []string{"-la"},
		},
		{
			Name:   "ls -l",
			Binary: ls,
			Args:   []string{"-la"},
		},
	}).Show()
	if err != nil {
		_ = fmt.Errorf(err.Error())
	}
	fmt.Println("selected: ",res.Selected().Name)
}
