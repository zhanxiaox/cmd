package cmd

import (
	"fmt"

	"github.com/zhanxiaox/cmd"
)

func init() {
	downloadCmd := cmd.Command{
		Name:        "download",
		Desc:        "download file form cloud",
		DefaultHelp: true,
		Flags: map[string]cmd.Flag{
			"-p": {Name: "process", Usage: "download threads(1-10)"},
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
