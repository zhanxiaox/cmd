package cmd

import "fmt"

func (this *app) addDefaultHelpCommand() {
	if this.defaultHelp {
		if _, ok := this.commands["help"]; !ok {
			this.AddCommand(Command{Name: "help", Desc: "default help for this app,you can overide this with addCommand", Excute: func(cmd *Command) {
				this.helpApp()
			}})
		}
	}
}

func (this *Command) addDefaultHelpFlag() {
	if this.DefaultHelp {
		if _, ok := this.Flags["-h"]; !ok {
			this.Flags["-h"] = Flag{Name: "help", Usage: "default help for this command", Excute: func(this *Command) {
				this.helpCommad()
			}}
		}
	}
}

func (this *app) helpApp() {
	fmt.Println(this.name, this.version)
	fmt.Println(this.desc)
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println(this.name, "[command] [flag]")
	fmt.Println()
	fmt.Println("Available Commands:")
	for _, command := range this.commands {
		fmt.Println(command.Name, generateSpace(20-len(command.Name)), command.Desc)
	}
	fmt.Println()
	fmt.Println("Available Flags:")
	for _, command := range this.commands {
		if len(command.Flags) > 0 {
			fmt.Println(command.Name)
			for k, fl := range command.Flags {
				fmt.Println(generateSpace(2), k, generateSpace(16-len(k)), fl.Usage)
			}
		}

	}
}

func (this *Command) helpCommad() {
	fmt.Println("Usage:")
	fmt.Println(args.Path, "[command] [flag]")
	fmt.Println()
	fmt.Println("Available Commands:")
	fmt.Println(this.Name, generateSpace(20-len(this.Name)), this.Desc)
	fmt.Println()
	fmt.Println("Available Flags:")
	if this.Flags != nil {
		fmt.Println(this.Name)
		for k, fl := range this.Flags {
			fmt.Println(generateSpace(2), k, generateSpace(16-len(k)), fl.Usage)
		}
	}
}
