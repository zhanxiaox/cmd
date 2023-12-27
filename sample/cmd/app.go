package cmd

import (
	"github.com/zhanxiaox/cmd"
)

var app = cmd.New()

func Excute() {

	// optional config,it will show in default help command
	app.Name = "fsync"
	app.Desc = "fsync is ..."
	app.Version = "0.0.1"

	// add help commands with default help
	app.AddCommand(cmd.Command{Name: "help", Desc: "help for this app", Excute: func(c cmd.Command) error {
		return app.DefaultHelp()
	}})
	if err := app.Excute(); err != nil {
		app.Warn(err.Error())
	}
}
