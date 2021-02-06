package action

import (
	"github.com/alexpfx/go_action/input"
	"log"
	
	"io"
	"os/exec"
	"strings"
)

type Runnable interface {
	Run() (output []byte, err error)
}

type Action struct {
	Name        string
	Binary      *Binary
	Args        []string
	InputConfig *input.Config
	output      []byte
	
	
	InputFromPipe bool
	Next          *Action
	Converter     func([]byte) ([]byte, error)
	prev          *Action
}



type Actual struct {
	action    *Action
	Pipe      bool
	Converter func([]byte) ([]byte, error)
}

func (c Actual) Action() *Action {
	return c.action
}
func (c Prev) Action() *Action {
	return c.action
}

type Prev struct {
	action    *Action
}


func (c Actual) Run() (output []byte, err error) {
	return c.action.Run()
}

func (c *Action) Run() (output []byte, err error) {
	var readInput []string
	
	inputConfig := c.InputConfig
	
	if inputConfig != nil {
		resolver := inputConfig.Resolver
		readInput, _ = resolver.Resolve(inputConfig)
	}
	
	var args []string
	
	args = append(args, c.Binary.FixArgs...)
	args = append(args, c.Args...)
	args = append(args, readInput...)
	
	command := exec.Command(
		c.Binary.CmdPath,
		args...,
	)
	
	convertAndAppend(command, c)
	
	c.output, err = command.CombinedOutput()
	if err != nil {
		return nil, err
	}
	
	if c.Next != nil {
		actual := c.Next
		actual.prev = c
		return actual.Run()
	}
	return c.output, err
}

func appendOutput(command *exec.Cmd, pipe bool, output []byte) {
	if output == nil {
		return
	}
	
	if pipe {
		stdin, err := command.StdinPipe()
		if err != nil {
			log.Fatal(err)
		}
		go func() {
			defer stdin.Close()
			_, _ = io.WriteString(stdin, string(output))
		}()
		
	} else {
		command.Args = append(command.Args, strings.TrimSpace(string(output)))
	}
}

func convertAndAppend(cmd *exec.Cmd, actualAction *Action) {
	prev := actualAction.prev
	
	if prev == nil {
		return
	}
	
	var prevOut []byte
	var err error
	
	if prev.Converter != nil {
		prevOut, err = prev.Converter(actualAction.output)
		checkError(err)
	} else {
		prevOut = prev.output
	}
	
	appendOutput(cmd, actualAction.InputFromPipe, prevOut)
	
}

type Pipe struct {
	Cmd Action
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
