package linux

import "github.com/alexpfx/go_action/action"

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
	
	
)
