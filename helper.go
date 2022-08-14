/* helpers for the Unix Command line interpreter */
package main

import (
	"fmt"
	"os"
)
/**
 * Command struct
 * name: name of the command
 * description: description of the command
 * action: action to be performed when the command is run
 */
type Command struct {
	Name        string
	Description string
	Action      func(args []string)
}

/**
 * commands is a list of Command structs
 */
var commands = []Command{
	{"cd", "change directory", cd},
	{"pwd", "print working directory", pwd},
	{"ls", "list files in current directory", ls},
	{"mkdir", "make directory", mkdir},
	{"rmdir", "remove directory", rmdir},
	{"rm", "remove file", rm},
	{"touch", "create empty file", touch},
	{"cat", "print file contents", cat},
	{"cp", "copy file", cp},
	{"mv", "move file", mv},
	{"exit", "exit the shell", exit},
}

/**
 * getCommand - get the command from the command line
 * @name: name of the command
 * return: pointer to the command struct
 */
func getCommand(name string) *Command {
	for _, command := range commands {
		if command.Name == name {
			return &command
		}
	}
	return nil
}


/**
 * runCommand - run the command
 * @command: string containing the command
 * @args: arguments to the command
 */

func runCommand(command string, args []string) {
	cmd := getCommand(command)
	if cmd == nil {
		fmt.Println("unknown command:", command)
		return
	}
	cmd.Action(args)
}

/**
 * prompt for input and run the command
 **/

func prompt() {
	fmt.Printf("$ ")
}

/** Colorize - colorize file names
 * @param s - file name
 * @return - colored file name
 * check if file is a directory or is an executable file
 * otherwise colorize the file name in white
 */

func Colorize(s string) string {

	fileInfo, err := os.Stat(s)
	if err != nil {
		return s
	}
	if fileInfo.IsDir() {
		return fmt.Sprintf("\x1b[34m%s\x1b[0m", s)
	}
	/* colorize regular files in white */
	if fileInfo.Mode().IsRegular() {
		return fmt.Sprintf("\x1b[37m%s\x1b[0m", s)
	}
	/* check if executable using bitmask 0100 */
	if fileInfo.Mode().Perm()&0100 != 0 {
		return fmt.Sprintf("\x1b[33m%s\x1b[0m", s)
	}
	return s
}