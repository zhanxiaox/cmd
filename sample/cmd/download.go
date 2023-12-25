package cmd

import (
	"fmt"

	"github.com/zhanxiaox/cmd"
)

func init() {
	downloadCmd := cmd.Command{
		Name:        "download",
		Desc:        "download basic file form vesdk",
		DefaultHelp: true,
		Flags: map[string]cmd.Flag{
			"-p": {Name: "process", Usage: "download threads(1-10)"},
		},
		Excute: func(this *cmd.Command) {
			p, err := this.MustGetFlagInt64("-p")
			if err != nil {
				fmt.Println("mustget error:", err.Error())
				return
			}
			fmt.Println(p)
		},
	}

	app.AddCommand(downloadCmd)
}
