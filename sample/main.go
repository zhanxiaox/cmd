package main

import (
	"github.com/zhanxiaox/cmd"
)

func main() {
	app := cmd.New("app", "1.0", "test app")

	download := cmd.Command{
		Name: "download",
		Desc: "download basic file",
		Flags: []cmd.Flag{
			{Name: "process", Default: 1, Usage: "并发下载数"},
		},
		Run: func(this *cmd.Command) {},
	}

	app.AddCommand(download)
	app.Excute()
}
