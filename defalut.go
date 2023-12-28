package cmd

import "fmt"

func (this app) DefaultHelp() error {
	fmt.Println(this.Name, this.Version)
	fmt.Println(this.Desc)
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println(args.Path, "[command] [flag]")
	fmt.Println()
	fmt.Println("Available Commands:")
	for _, command := range this.commands {
		fmt.Println(command.Name, generateSpace(20-len(command.Name)), command.Desc)
	}
	fmt.Println()
	fmt.Println(`Use "` + args.Path + ` [command] --help" for more information about a command.`)
	// fmt.Println("Available Flags:")
	// for _, command := range this.commands {
	// 	if len(command.Flags) > 0 {
	// 		fmt.Println(command.Name)
	// 		for k, fl := range command.Flags {
	// 			fmt.Println(generateSpace(2), k, generateSpace(16-len(k)), fl.Usage)
	// 		}
	// 	}
	// }
	return nil
}

func (this Command) DefaultHelp() error {
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
	return nil
}
