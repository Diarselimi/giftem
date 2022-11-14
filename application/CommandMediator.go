package application

import "sync"

type Command interface {
	Execute()
}

type CommandMediator struct {
	Mu       *sync.Mutex
	commands []Command
}

func (cm *CommandMediator) Add(cmd Command) {
	cm.commands = append(cm.commands, cmd)
}

func (cm *CommandMediator) Run() {
	cm.Mu.Lock()
	defer cm.Mu.Unlock()

	for _, command := range cm.commands {
		command.Execute()
	}
}
