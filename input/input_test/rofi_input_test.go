package action_test

import (
	"github.com/alexpfx/go_action/input"
	"github.com/stretchr/testify/assert"
	"testing"
)
//FIXME deixar teste estático e não interativo
func TestRofiResolver_Resolve(t *testing.T) {
	resolver := input.RofiResolver{}
	config := &input.ResolverConfig{
		Resolver: resolver,
		Keys: []string{"argkey"},
	}
	
	res, _ := resolver.Resolve(config)
	
	assert.Equal(t, []string{"argkey", "argvalue"}, res)
	
}

//FIXME deixar teste estático e não interativo
func TestRofiResolver_ResolveArgSep(t *testing.T) {
	resolver := input.RofiResolver{}
	config := &input.ResolverConfig{
		Resolver: resolver,
		Keys: []string{"argkey"},
		ArgSep: "\n",
	}
	
	res, _ := resolver.Resolve(config)
	
	assert.Equal(t, []string{"argkey", "argvalue"}, res)
	
}
