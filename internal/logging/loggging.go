package logging

import (
	"fmt"
)

func Verbose(verbose bool, format string, a ...any) {
	if !verbose {
		return
	}

	fmt.Printf("[LOG]: "+format+"\n", a...)
}
