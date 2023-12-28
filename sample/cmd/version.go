package cmd

import (
	"fmt"

	"github.com/zhanxiaox/cmd"
)

func init() {
	versionCmd := cmd.Command{
		Name: "version",
		Desc: "Show app version",
		Flags: map[string]cmd.Flag{
			"-h": {Name: "help", Usage: "Print this message", Excute: func(this cmd.Command) error {
				fmt.Println(this.Flags["-h"].Usage)
				app.Info(this.Flags["-h"].Usage)
				return nil
			}},
		},
		Excute: func(this cmd.Command) error {
			app.Info(app.Version)
			return nil
		},
	}

	app.AddCommand(versionCmd)
}
