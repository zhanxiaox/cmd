package cmd

// 设置空格位
func generateSpace(i int) string {
	space := ""
	for i >= 0 {
		space += " "
		i--
	}
	return space
}
