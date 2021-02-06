package action

import (
	"fmt"
	"github.com/alexpfx/go_menus/menu"
	"github.com/alexpfx/go_menus/menu/fzf"
	"log"
	"strconv"
	"strings"
)

const sep = ";"

func NewFzfTree(prompt string, actions []Action) Tree {
	
	fzfMenu := fzf.NewBuilder().Prompt("selecione").
		AutoSelect(true).
		Prompt(prompt).(fzf.Builder).
		WithNth("2", sep).Build()
	
	return tree{
		menu:    fzfMenu,
		actions: actions,
	}
	
}

type Tree interface {
	Show() (Result, error)
}

type Result interface {
	Out() string
	Selected() *Action
}

type result struct {
	out      string
	selected *Action
}

func (r result) Out() string {
	return r.out
}

func (r result) Selected() *Action {
	return r.selected
}

type tree struct {
	menu    menu.Menu
	actions []Action
}

func (t tree) Show() (Result, error) {
	var sb strings.Builder
	for i, action := range t.actions {
		sb.WriteString(fmt.Sprintf("%d;%s\n", i, action.Name))
	}
	
	out, err := t.menu.Run(sb.String())
	if err != nil {
		return nil, err
	}
	var selected *Action
	selected = getSelected(t, out)
	
	
	return result{
		selected: selected,
		out:      out,
		
	}, err
	
}

func getSelected(t tree, out string) *Action {
	if len(t.actions) == 0 || out == "" {
		return nil
	}
	strIndex := strings.Split(out, sep)
	if len(strIndex) == 0 {
		return nil
	}
	index, err := strconv.Atoi(strIndex[0])
	
	if err != nil {
		log.Fatal(err)
	}
	if index > len(t.actions) {
		log.Fatalf("erro ao obter selecionado: indice: %d size: %d", index, len(t.actions))
	}
	
	return &t.actions[index]
}
