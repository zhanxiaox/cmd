package cmd

import (
	"errors"
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
	flags map[string]Flag
}

type Flag struct {
	Name  string
	value string
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

	call := false
	call_command := args[0]
	call_flag := args[1:]
	if len(call_flag)%2 != 0 {
		fmt.Println("flag param error,param must be key=value pair")
		return
	}

	for _, command := range this.commands {
		call = false
		if command.Name == call_command {
			call = true
			for k, v := range call_flag {
				if k%2 == 1 {
					continue
				}

				input_flag_value := call_flag[k+1]

				if a, ok := command.flags[v]; ok {
					a.value = input_flag_value
					command.flags[v] = a
				}
			}
			command.Run(&command)
		}
	}

	if call == false {
		fmt.Println("command", call_command, "not found")
	}
}

func (this *Command) AddFlag(f Flag) {
	if this.flags == nil {
		this.flags = map[string]Flag{f.Name: f}
	} else {
		this.flags[f.Name] = f
	}
}

func (this *Command) mustGet(k string) (string, error) {
	v, ok := this.flags[k]
	if !ok {
		return "", errors.New("flag" + k + "not found")
	}
	return v.value, nil
}

func (this *Command) shouldGet(k string) string {
	v, _ := this.flags[k]
	return v.value
}

func (this *Command) MustGetFlagInt64(k string) (int64, error) {
	str, err := this.mustGet(k)
	if err != nil {
		return 0, err
	}
	i64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return i64, nil
}

func (this *Command) MustGetFlagString(k string) (string, error) {
	return this.mustGet(k)
}

func (this *Command) MustGetFlagBool(k string) (bool, error) {
	str, err := this.mustGet(k)
	if err != nil {
		return false, err
	}
	b, err := strconv.ParseBool(str)
	if err != nil {
		return false, err
	}
	return b, nil
}

func (this *Command) ShouldGetFlagInt64(k string) int64 {
	str := this.shouldGet(k)
	i64, _ := strconv.ParseInt(str, 10, 64)
	return i64
}

func (this *Command) ShouldGetFlagString(k string) string {
	return this.shouldGet(k)
}

func (this *Command) ShouldGetFlagBool(k string) bool {
	str := this.shouldGet(k)
	b, _ := strconv.ParseBool(str)
	return b
}

func New(name, version, desc string) *App {
	return &App{Name: name, Version: version, Desc: desc}
}
