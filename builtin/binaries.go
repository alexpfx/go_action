package builtin

import (
	"github.com/alexpfx/go_action/action"
)

var (
	Bat = &action.Binary{
		CmdPath: "/usr/bin/bat",
		Name:    "bat",
	}
	
	Echo = &action.Binary{
		CmdPath: "echo",
		Name:    "echo -e",
		FixArgs: []string{"-e"},
	}
	Sort = &action.Binary{
		CmdPath: "sort",
		Name:    "sort",
	}
	
	Ls = &action.Binary{
		CmdPath: "ls",
		Name:    "ls",
		FixArgs: nil,
	}
	
	Jq = &action.Binary{
		CmdPath: "jq",
		Name:    "jq JSON processor",
		FixArgs: []string{"--tab"},
	}
	
	Xsel = &action.Binary{
		CmdPath: "xsel",
		Name:    "xel",
		FixArgs: []string{"-b"},
	}
)
