package cmd

import "github.com/zhanxiaox/cmd"

var app = cmd.New("fsync", "1.0.0", "test", true)

func Excute() {
	app.Excute()
}
