package cmd

import (
	"github.com/zhanxiaox/cmd"
)

func init() {
	uploadCmd := cmd.Command{
		Name: "upload",
		Desc: "upload file to cloud",
		Excute: func(this cmd.Command) error {
			return nil
		},
	}

	app.AddCommand(uploadCmd)
}
