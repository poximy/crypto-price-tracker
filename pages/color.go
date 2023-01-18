package pages

import "fmt"

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

var Bold = "\033[1m"
var Dim = "\033[2m"
var Italic = "\033[3m"

var Clear = "\033[H\033[2J"

func MakePrint(prefix, suffix string) (Print func(...any), Println func(...any)) {
	return func(a ...any) {
			fmt.Print(prefix)
			fmt.Print(a...)
			fmt.Print(suffix)
		}, func(a ...any) {
			fmt.Print(prefix)
			fmt.Print(a...)
			fmt.Println(suffix)
		}
}
