package builtin

import (
	"bufio"
	"github.com/alexpfx/go_action/input"
	"os"
)

type LinuxReadInput struct {
}

func (f LinuxReadInput) Resolve(config *input.ResolverConfig) ([]string, error) {
	read := make([]string, 0)
	
	var scanner *bufio.Scanner
	
	if config.Reader != nil {
		scanner = bufio.NewScanner(config.Reader)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}
	
	for {
		scanner.Scan()
		
		text := scanner.Text()
		
		if len(text) != 0 {
			read = append(read, text)
		} else {
			break
		}
		
	}
	res := make([]string, 0)
	
	for i, s := range read {
		if i < len(config.Keys) {
			res = append(res, config.Keys[i])
		}
		res = append(res, s)
	}
	return res, nil
}
