package utils

import "fmt"

const (
	Reset = "\033[0m"
	Bold  = "\033[1m"
	Under = "\033[4m"
)

// Text colors
const (
	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
)

// Background colors
const (
	BgBlack   = "\033[40m"
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"
)

func ColorText(color string, text string, styles ...string) string {
	stylePrefix := ""
	for _, s := range styles {
		stylePrefix += s
	}
	return fmt.Sprintf("%s%s%s", stylePrefix+color, text, Reset)
}
