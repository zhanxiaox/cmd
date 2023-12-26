package cmd

import (
	"github.com/zhanxiaox/cmd"
)

var app = cmd.New(cmd.App{

	// APP name
	Name: "fsync",

	// APP version
	Version: "0.0.1",

	// APP description
	Desc: "fsync is ...",

	// APP whether to automatically add help commands
	DefaultHelp: true,
})

func Excute() {
	if err := app.Excute(); err != nil {
		app.Warn(err.Error())
	}
}
