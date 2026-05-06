package reader

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			fmt.Println()
			fmt.Fprintln(os.Stderr, "Aborted.")
			os.Exit(0)
		} else {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			os.Exit(1)
		}
	}
	return strings.TrimSpace(text)
}

func ConfirmAction(prompt string, defaultYes bool) bool {
	yn := " [y/N]: "
	if defaultYes {
		yn = " [Y/n]: "
	}
	res := ReadInput(prompt + yn)
	if res == "" {
		return defaultYes
	}
	return strings.EqualFold(res, "y")
}
