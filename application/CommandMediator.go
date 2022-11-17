package application

import (
	"fmt"
	"giftem/entity"
	"net/http"
	"sync"
)

type Command interface {
	Execute() entity.Gift
}

type CommandMediator struct {
	Mu       *sync.Mutex
	commands []Command
}

func (cm *CommandMediator) Add(cmd Command) {
	cm.commands = append(cm.commands, cmd)
}

func (cm *CommandMediator) Run(w *http.ResponseWriter) {
	cm.Mu.Lock()
	defer cm.Mu.Unlock()

	for _, command := range cm.commands {
		gf := command.Execute()
		fmt.Fprintf(*w, "<h1> Your gift is, %s </h1>", gf.Name)
	}
}
