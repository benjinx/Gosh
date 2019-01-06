package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	builtin = map[string]func([]string){}
	cwd     = "./"
)

func init() {
	builtin["exit"] = cmdExit
	builtin["pwd"] = cmdPwd
	builtin["echo"] = cmdEcho
}

func main() {

	// Gets the directory passed in or uses default
	if len(os.Args) > 1 {
		cwd = os.Args[1]
	} else {
		cwd, _ = os.Getwd()
	}

	// Hello Program
	fmt.Println("Oh Gosh!")

	// Initial Prompt
	fmt.Print("> ")

	// Scan the file
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		// Store the prompt that was input
		line := process(scanner.Text())

		// Check if there is any text
		if line == "" {
			continue
		}

		// Parse the string
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}

		// Compare against map of commands and target the correct command
		if cmd, found := builtin[args[0]]; found {
			cmd(args)
		} else {
			fmt.Fprintf(os.Stderr, "Command not found '%v'\n", line)
		}

		// Need to reprint Prompt for the next input
		fmt.Print("> ")
	}

}

func cmdExit([]string) {
	os.Exit(0)
}

func cmdPwd([]string) {
	abs, _ := filepath.Abs(cwd)
	fmt.Println(abs)
}

func cmdEcho(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

func process(line string) string {
	re := regexp.MustCompile("\\$([a-zA-Z_]*)")
	line = re.ReplaceAllStringFunc(line, func(name string) string {
		return os.Getenv(name[1:])
	})
	return line
}
