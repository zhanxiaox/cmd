package cmd

import "fmt"

// 生成空格
func generateSpace(i int) string {
	space := ""
	for i >= 0 {
		space += " "
		i--
	}
	return space
}

func (this app) Info(s string) {
	fmt.Println("[INFO]:", s)
}

func (this app) Warn(s string) {
	fmt.Println("[WARN]:", s)
}
