package binary

var (
	Bat = &Binary{
		CmdPath: "/usr/bin/bat",
		Name:    "bat",
	}
	
	Echo = &Binary{
		CmdPath: "echo",
		Name:    "echo -e",
		FixArgs: []string{"-e"},
	}
	Sort = &Binary{
		CmdPath: "sort",
		Name:    "sort",
	}
	
	Ls = &Binary{
		CmdPath: "ls",
		Name:    "ls",
		FixArgs: nil,
	}
	
	Jq = &Binary{
		CmdPath: "jq",
		Name:    "jq JSON processor",
		FixArgs: []string{"--tab"},
	}
	
	Xsel = &Binary{
		CmdPath: "xsel",
		Name:    "xel",
		FixArgs: []string{"-b"},
	}
	
	
)
