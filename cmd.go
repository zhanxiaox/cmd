package cmd

import (
	"flag"
	"fmt"
	"strconv"
)

type App struct {
	Name     string
	Version  string
	Desc     string
	commands []Command
}

type Command struct {
	Name  string
	Desc  string
	Help  func(*Command)
	Run   func(*Command)
	Flags map[string]Flag
}

type Flag struct {
	Name  string
	Value any
	Usage string
}

func (this *App) Help() {
	fmt.Println(this.Name, this.Version)
	fmt.Println(this.Desc)
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println(this.Name, "[command] [flag]")
	fmt.Println("")
	fmt.Println("Available Commands:")
	fmt.Println("download some balaaaa...")
	fmt.Println("upload   some balaaaaaaa...")
	fmt.Println("")
	fmt.Println("Flags:")
	fmt.Println("-h balaaaaaa...")
}

func (this *App) AddCommand(command Command) {
	this.commands = append(this.commands, command)
}

func (this *App) Excute() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		this.Help()
		return
	}

	input_command := args[0]
	input_flag := args[1:]
	find := false
	if len(input_flag)%2 != 0 {
		fmt.Println("flag param error,param must be key=value pair")
		return
	}
	for _, command := range this.commands {
		if command.Name == input_command {
			find = true
			for k, v := range input_flag {
				if k%2 == 1 {
					continue
				}
				if _, ok := command.Flags[v]; ok {
					a := command.Flags[v]
					a.Value = input_flag[k+1]
					command.Flags[v] = a
				}
			}
			command.Run(&command)
		}
	}

	if find == false {
		fmt.Println("command", input_command, "not found")
	}
}

func (this *Command) GetFlagInt64(k string) int64 {
	v := this.Flags[k]
	fmt.Println(v, v.Value)
	a, _ := v.Value.(string)
	c, _ := strconv.ParseInt(a, 10, 64)
	// fmt.Println(a, ok)
	return c
}

func New(name, version, desc string) *App {
	return &App{Name: name, Version: version, Desc: desc}
}
