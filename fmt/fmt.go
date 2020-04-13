package xfmt

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

//彩色打印文字
func PrintMsg(msg string, color int) {
	fmt.Printf("%c[%d;%d;%dm%s%c[0m\n", 0x1B, 0, 0, color, msg, 0x1B)
}
