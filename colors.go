// source --> https://github.com/alexeyco/simpletable/blob/v1/_example/05-colorize/main.go

package tasky

import "fmt"

const (
	ColorDefault = "\x1b[39m"

	ColorRed   = "\x1b[91m"
	ColorGreen = "\x1b[32m"
	ColorBlue  = "\x1b[94m"
	ColorGray  = "\x1b[90m"
)

func Red(s string) string {
	return fmt.Sprintf("%s%s%s", ColorRed, s, ColorDefault)
}

func Green(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGreen, s, ColorDefault)
}

func Blue(s string) string {
	return fmt.Sprintf("%s%s%s", ColorBlue, s, ColorDefault)
}

// func gray(s string) string {
// 	return fmt.Sprintf("%s%s%s", ColorGray, s, ColorDefault)
// }
