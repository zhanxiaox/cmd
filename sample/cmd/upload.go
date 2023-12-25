package cmd

import (
	"github.com/zhanxiaox/cmd"
)

func init() {
	uploadCmd := cmd.Command{
		Name: "upload",
		Desc: "upload file to oss",
		Flags: map[string]cmd.Flag{
			"--process": {Name: "process", Usage: "upload threads(1-10)"},
			// "-h": {Name: "help", Usage: "help message", Executable: true, Excute: func(this *cmd.Command) {
			// 	fmt.Println("hellp upload world")
			// }},
		},
		Excute: func(this *cmd.Command) {
			// p, err := this.MustGetFlagInt64("--process")
			// if err != nil {
			// 	fmt.Println(err.Error())
			// 	return
			// }
			// fmt.Println(p)
		},
	}

	app.AddCommand(uploadCmd)
}
