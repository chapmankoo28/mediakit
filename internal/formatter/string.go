package formatter

import "fmt"

func Center(s string, width int) string {
	pad := (width - len(s)) / 2
	return fmt.Sprintf("%*s", pad+len(s), s)
}
