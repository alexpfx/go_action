package action

import (
	"fmt"
	"github.com/alexpfx/go_action/internal/util"
	"github.com/alexpfx/go_menus/menu"
	"log"
	"strconv"
	"strings"
)

const sep = ";"


type Tree interface {
	Show() (*Action, bool, error)
}


type tree struct {
	menu    menu.Menu
	actions []Action
}

func (t tree) Show() (*Action, bool, error) {
	var sb strings.Builder
	for i, action := range t.actions {
		sb.WriteString(fmt.Sprintf("%d;%s\n", i, action.Name))
	}
	
	out, err := t.menu.Run(sb.String())
	if err != nil {
		return nil, false, err
	}
	
	selected, found := getSelected(t, out)
	
	return selected, found, err
	
}

func getSelected(t tree, out string) (*Action, bool) {
	if len(t.actions) == 0 || out == "" {
		return nil, false
	}
	strIndex := util.SplitSep(out, sep)
	if len(strIndex) == 0 {
		return nil, false
	}
	index, err := strconv.Atoi(strIndex[0])
	if err != nil {
		return nil, false
	}
	if index > len(t.actions) {
		log.Fatalf("erro ao obter selecionado: indice: %d size: %d", index, len(t.actions))
	}
	
	return &t.actions[index], true
}
