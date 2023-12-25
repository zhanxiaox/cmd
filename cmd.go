package cmd

import (
	"fmt"
)

type app struct {
	name        string
	version     string
	desc        string
	commands    map[string]Command
	defaultHelp bool
}

type Command struct {
	Name        string
	Desc        string
	Excute      func(*Command)
	Flags       map[string]Flag
	DefaultHelp bool
}

type Flag struct {
	Name   string
	value  string
	Usage  string
	Excute func(*Command)
}

func (this *app) AddCommand(command Command) {
	if command.Flags == nil {
		command.Flags = map[string]Flag{}
	}
	command.addDefaultHelpFlag()
	this.commands[command.Name] = command
}

var args = getArgs()

func (this *app) Excute() {
	if args.Command == "" {
		if command, ok := this.commands["help"]; ok {
			command.Excute(&command)
		} else {
			fmt.Println(this.name, "has not set help commands,if you want a default help,you can set cmd.New(*,*,*,true) to enable it or use addCommand to custom a help command")
		}
		return
	}

	if command, ok := this.commands[args.Command]; ok {
		flExcutable := false
		for k, v := range args.Flags {
			if fl, ok := command.Flags[k]; ok {
				fl.value = v
				command.Flags[k] = fl
				if fl.Excute != nil {
					flExcutable = true
					fl.Excute(&command)
					break
				}
			}
		}
		if !flExcutable {
			command.Excute(&command)
		}
	} else {
		fmt.Println("command", args.Command, "not found,run", this.name, "help to get more infomention")
	}
}

func New(name, version, desc string, defaultHelp bool) *app {
	instace := &app{name: name, version: version, desc: desc, defaultHelp: defaultHelp, commands: map[string]Command{}}
	instace.addDefaultHelpCommand()
	return instace
}
