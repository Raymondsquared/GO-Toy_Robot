package commands

import (
)

type Invoker struct {
    command Command
}

func NewInvoker(command Command) *Invoker {
	return &Invoker{ command }
} 

func(this *Invoker) Invoke() {
    this.command.Execute()
}