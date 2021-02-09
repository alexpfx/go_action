package test

import (
	"fmt"
	"github.com/alexpfx/go_action/action"
	"github.com/alexpfx/go_action/action/binary"
	
	"testing"
)

func echoAndSort(input string) *action.Action {
	return &action.Action{
		Name:   "echo | sort ",
		Binary: binary.Echo,
		
		Next: &action.Action{
			Binary:        binary.Sort,
			InputFromPipe: true,
		},
		Args: []string{
			input,
		},
		
	}
}

func echoAndLs(input string) *action.Action {
	return &action.Action{
		Name:   "echo > ls",
		Binary: binary.Echo,
		Args:   []string{input},
		Next: &action.Action{
			Binary: binary.Ls,
		},
	}
}

func TestRunPipeToNext(t *testing.T) {
	x := echoAndSort("palmeiras\nvasco\ngremio\n")
	res, err := x.Run()
	fmt.Println("res: ", string(res), "err: ", err)
}

func TestRunAppendArgsToNext(t *testing.T) {
	x := echoAndLs("-la")
	
	res, err := x.Run()
	
	fmt.Println("res: ", string(res), "err: ", err)
}
