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

					switch a.Value.(type) {
					case int64:

					case string:

					case bool:
					default:
					}

					a.Value = input_flag_value
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
	a, ok := v.Value.(string)
	if !ok {
		return "", errors.New("flag value error,not string")
	}
	return a, nil
}

func (this *Command) shouldGet(k string) any {
	v, _ := this.flags[k]
	return v.Value
}

func (this *Command) MustGetFlagInt64(k string) (int64, error) {
	v, err := this.mustGet(k)
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (this *Command) MustGetFlagString(k string) (string, error) {
	return this.mustGet(k)
}

func (this *Command) MustGetFlagBool(k string) (bool, error) {
	v, err := this.mustGet(k)
	if err != nil {
		return false, err
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		return false, err
	}
	return b, nil
}

func (this *Command) ShouldGetFlagInt64(k string) uint64 {
	a := this.shouldGet(k)
	i64, ok := a.(uint64)
	fmt.Println(i64)
	if ok {
		return i64
	}
	str, _ := a.(string)
	i64, _ = strconv.ParseUint(str, 10, 64)
	return i64
}

func (this *Command) ShouldGetFlagString(k string) string {
	a := this.shouldGet(k)
	return a.(string)
}

func (this *Command) ShouldGetFlagBool(k string) bool {
	a := this.shouldGet(k)
	str, _ := a.(string)
	b, _ := strconv.ParseBool(str)
	return b
}

func New(name, version, desc string) *App {
	return &App{Name: name, Version: version, Desc: desc}
}
