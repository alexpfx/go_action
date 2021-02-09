package action_test

import (
	"github.com/alexpfx/go_action/input"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_LinuxRead(t *testing.T) {
	resolver := input.ScannerResolver{}
	var config *input.ResolverConfig
	
	config = &input.ResolverConfig{
		Resolver: resolver,
		Keys:     []string{"argkey"},
		ArgSep:   "\n",
		Reader:   strings.NewReader("argvalue\n"),
	}
	
	res, _ := resolver.Resolve(config)
	
	assert.Equal(t, []string{"argkey", "argvalue"}, res)
	
}
