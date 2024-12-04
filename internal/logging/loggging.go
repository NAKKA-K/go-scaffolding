package logging

import (
	"fmt"
)

func Verbose(verbose bool, a ...any) {
	if !verbose {
		return
	}

	fmt.Print("[LOG]: ")
	fmt.Println(a...)
}
