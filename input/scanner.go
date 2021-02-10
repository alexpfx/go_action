package input

import (
	"bufio"
	"fmt"
	"os"
)

type ScannerResolver struct {
}

func (f ScannerResolver) Resolve(config *ResolverConfig) ([]string, error) {
	read := make([]string, 0)
	
	var scanner *bufio.Scanner
	
	if config.Reader != nil {
		scanner = bufio.NewScanner(config.Reader)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}
	
	fmt.Printf("%s: ", config.Prompt)
	scanner.Scan()
	
	text := scanner.Text()
	
	if len(text) != 0 {
		read = append(read, text)
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
