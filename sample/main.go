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
		Run: func(this *cmd.Command) {
			fmt.Println(this.ShouldGetFlagInt64("--process"))
			p, err := this.MustGetFlagInt64("--process")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(p)
		},
	}
	download.AddFlag(cmd.Flag{Name: "--process", Value: uint64(1), Usage: "并发下载数"})
	download.AddFlag(cmd.Flag{Name: "-p", Value: 1, Usage: "并发下载数"})
	app.AddCommand(download)
	app.Excute()
}
