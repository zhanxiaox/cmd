package cmd

import (
	"fmt"

	"github.com/zhanxiaox/cmd"
)

func init() {
	downloadCmd := cmd.Command{
		Name: "download",
		Desc: "Download file from cloud",
		Flags: map[string]cmd.Flag{
			"-p": {Name: "process", Usage: "Download threads(1-10)"},
			"-h": {Name: "help", Usage: "Help this command", Excute: func(this cmd.Command) error {
				return this.DefaultHelp()
			}},
		},
		Excute: func(this cmd.Command) error {
			p, err := this.MustGetFlagInt64("-p")
			if err != nil {
				return err
			}
			app.Info("start download files with " + fmt.Sprint(p) + " threads")
			return nil
		},
	}

	app.AddCommand(downloadCmd)
}
