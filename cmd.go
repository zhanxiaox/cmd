package cmd

import (
	"errors"
	"fmt"
)

type app struct {
	Name     string
	Version  string
	Desc     string
	commands map[string]Command
}

type Command struct {
	Name   string
	Desc   string
	Excute func(Command) error
	Flags  map[string]Flag
}

type Flag struct {
	Name   string
	value  string
	Usage  string
	Excute func(Command) error
}

func (this *app) AddCommand(command Command) {
	if this.commands == nil {
		this.commands = map[string]Command{}
	}
	if command.Flags == nil {
		command.Flags = map[string]Flag{}
	}
	this.commands[command.Name] = command
}

var args = getArgs()

func (this app) Excute() error {
	if args.Command == "" {
		if command, ok := this.commands["help"]; ok {
			return command.Excute(command)
		} else {
			return errors.New(fmt.Sprint(this.Name, " has not set help command or default help command,if you want a default help,you can set cmd.New(cmd.App{DefaultHelp:true}) to enable it or use addCommand to custom a help command"))
		}
	}
	if command, ok := this.commands[args.Command]; ok {
		for k, v := range args.Flags {
			if fl, ok := command.Flags[k]; ok {
				fl.value = v
				command.Flags[k] = fl
				if fl.Excute != nil {
					return fl.Excute(command)
				}
			}
		}
		return command.Excute(command)
	} else {
		return errors.New(fmt.Sprint("command ", args.Command, " not found,run ", this.Name, " help to get more infomention"))
	}
}

func New() *app {
	return &app{}
}
