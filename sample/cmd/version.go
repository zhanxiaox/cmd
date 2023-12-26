package cmd

import (
	"github.com/zhanxiaox/cmd"
)

func init() {
	versionCmd := cmd.Command{
		Name:        "version",
		Desc:        "show app version",
		DefaultHelp: false,
		Excute: func(this cmd.Command) error {
			app.Info(app.Version)
			return nil
		},
	}

	app.AddCommand(versionCmd)
}
