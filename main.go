package main

import (
	"os"
	"bufio"
	"strings"
)


func main() {
	for {
		prompt()
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		text := input.Text()
		args := strings.Split(text, " ")
		runCommand(args[0], args[1:])
	}
}