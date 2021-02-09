package binaries

import "github.com/alexpfx/go_action/action/binary"

var (
	Bat = &binary.Binary{
		CmdPath: "/usr/bin/bat",
		Name:    "bat",
	}
	
	Echo = &binary.Binary{
		CmdPath: "echo",
		Name:    "echo -e",
		FixArgs: []string{"-e"},
	}
	Sort = &binary.Binary{
		CmdPath: "sort",
		Name:    "sort",
	}
	
	Ls = &binary.Binary{
		CmdPath: "ls",
		Name:    "ls",
		FixArgs: nil,
	}
	
	Jq = &binary.Binary{
		CmdPath: "jq",
		Name:    "jq JSON processor",
		FixArgs: []string{"--tab"},
	}
	
	Xsel = &binary.Binary{
		CmdPath: "xsel",
		Name:    "xel",
		FixArgs: []string{"-b"},
	}
	

	
	
)
