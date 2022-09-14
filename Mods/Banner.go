package Z

import "fmt"

var banner string = `
_____   __                 __ 
|     |_|__|.-----.-----.--|  |
|       |  ||     |  -__|  _  |
|_______|__||__|__|_____|_____|
`

func (Struct *Data) Banner() {
	fmt.Println("\033[31m", banner, "\033[39m")
	fmt.Println("\n")
}

/* Export with a capital leter or the following error will occure

# command-line-arguments
./main.go:37:3: D.banner undefined (cannot refer to unexported field or method banner)



func (Struct *Data) banner() {
	fmt.Println("\033[31m", banner, "\033[39m")
	fmt.Println("\n")
}

*/
