package builtin

import (
	"github.com/alexpfx/go_action/input"
	"github.com/atotto/clipboard"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ClipResolver(t *testing.T) {
	resolver := ClipResolver{}
	var config *input.ResolverConfig
	
	config = &input.ResolverConfig{
		Resolver: resolver,
		Keys:     []string{"arg"},
		ArgSep:   "\n",
	}
	_ = clipboard.WriteAll("argvalue")
	res, _ := resolver.Resolve(config)
	assert.Equal(t, []string{"arg", "argvalue"}, res)
	
	_ = clipboard.WriteAll("myargvalue")
	config2 := &input.ResolverConfig{
		Resolver: resolver,
		Keys:     []string{"myarg"},
		ArgSep:   "\n",
	}
	res, _ = resolver.Resolve(config2)
	assert.Equal(t, []string{"myarg", "myargvalue"}, res)
	
	_ = clipboard.WriteAll("v1|v2")
	config3 := &input.ResolverConfig{
		Resolver: resolver,
		Keys:     []string{"myarg"},
		ArgSep:   "|",
	}
	res, _ = resolver.Resolve(config3)
	assert.Equal(t, []string{"myarg", "v1", "v2"}, res)
}
