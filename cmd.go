package cmd

import (
	"fmt"
)

type App struct {
	Name        string
	Version     string
	Desc        string
	commands    map[string]Command
	DefaultHelp bool
}

type Command struct {
	Name        string
	Desc        string
	Excute      func(*Command)
	Flags       map[string]Flag
	DefaultHelp bool
}

type Flag struct {
	Name       string
	value      string
	Usage      string
	Executable bool
	Excute     func(*Command)
}

func (this *App) AddCommand(command Command) {
	if command.Flags == nil {
		command.Flags = map[string]Flag{}
	}
	command.addDefaultHelpFlag()
	this.commands[command.Name] = command
}

var args = getArgs()

func (this *App) Excute() {
	if args.Command == "" {
		if command, ok := this.commands["help"]; ok {
			command.Excute(&command)
		} else {
			fmt.Println(this.Name, "has not set help commands,if you want a default help,you can set cmd.New(*,*,*,true) to enable it or use addCommand to custom a help command")
		}
		return
	}

	if command, ok := this.commands[args.Command]; ok {
		flExcutable := false
		for k, v := range args.Flags {
			if fl, ok := command.Flags[k]; ok {
				fl.value = v
				command.Flags[k] = fl
				if fl.Executable {
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
		fmt.Println("command", args.Command, "not found,run", this.Name, "help to get more infomention")
	}
}

func New(name, version, desc string, defaultHelp bool) *App {
	app := &App{Name: name, Version: version, Desc: desc, DefaultHelp: defaultHelp, commands: map[string]Command{}}
	app.addDefaultHelpCommand()
	return app
}
