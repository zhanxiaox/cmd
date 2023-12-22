package main

import (
	"fmt"

	"github.com/zhanxiaox/cmd"
)

func main() {
	app := cmd.New("fsync", "1.0.0", "test app")

	download := cmd.Command{
		Name: "download",
		Desc: "download basic file",
		Flags: map[string]cmd.Flag{
			"--process": {Name: "process", Value: 1, Usage: "并发下载数"},
		},
		Run: func(this *cmd.Command) {
			p := this.GetFlagInt64("--process")
			fmt.Println(p)
		},
	}
	app.AddCommand(download)
	app.Excute()
}
