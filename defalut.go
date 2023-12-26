package cmd

import "fmt"

func (this App) addDefaultHelpCommand() {
	this.AddCommand(Command{
		Name: "help",
		Desc: "default help for this app,you can override this with app.AddCommand(cmd.Command{Name:'help',...})",
		Excute: func(cmd Command) error {
			this.defaultHelpApp()
			return nil
		},
	})
}

func (this Command) addDefaultHelpFlag() {
	this.Flags["-h"] = Flag{
		Name:  "help",
		Usage: "default help for this command,you can override this with cmd.Command{Flags:map[string]cmd.Flag{Name:'-h',...}}",
		Excute: func(this Command) error {
			this.defaultHelpCommad()
			return nil
		},
	}
}

func (this App) defaultHelpApp() {
	fmt.Println(this.Name, this.Version)
	fmt.Println(this.Desc)
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println(this.Name, "[command] [flag]")
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

func (this Command) defaultHelpCommad() {
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
