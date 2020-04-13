package xfmt

import "fmt"

//颜色声明
const (
	Black = 30 + iota
	Red
	Green
	Yellow
	Blue
	Purple
	Cyan
	White
)

//PrintMsg 彩色打印文字
func PrintMsg(msg string, color int) {
	fmt.Printf("%c[%d;%d;%dm%s%c[0m", 0x1B, 0, 0, color, msg, 0x1B)
}

//PrintlnMsg 彩色打印文字，自动换行
func PrintlnMsg(msg string, color int) {
	PrintMsg(msg+"\n", color)
}
