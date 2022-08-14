package main

import (
	"fmt"
	"os"
	"path/filepath"
)

/**
 * cd - change directory
 * @param args - the arguments to the command
 */

func cd(args []string) {
	if len(args) != 1 {
		fmt.Println("usage: cd <dir>")
		return
	}
	dir := args[0]
	err := os.Chdir(dir)
	if err != nil {
		fmt.Println(err)
	}
}

/**
 * pwd - print working directory
 * @param args - the arguments to the command
 */

func pwd(s []string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dir)
}

/**
 * ls - list files in current directory
 * @param args - the arguments to the command
 * if no arguments are passed, list the current directory
 * if one argument is passed, list the contents of the directory
 */

func ls(args []string) {
	/* list files in current directory */
	files, err := filepath.Glob("*")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		/* colorize file names */
		fmt.Println(Colorize(file))
	}
	/* list files in specified directory */
	if len(args) == 1 {
		dir := args[0]
		files, err := filepath.Glob(dir + "/*")
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, file := range files {
			/* colorize file names */
			fmt.Println(Colorize(file))
		}
	}
}

/**
 * mkdir - make directory
 * @param args - the arguments to the command
 */
func mkdir(args []string) {
	if len(args) != 1 {
		fmt.Println("usage: mkdir <dir>")
		return
	}
	dir := args[0]
	err := os.Mkdir(dir, 0755)
	if err != nil {
		fmt.Println(err)
	}
}

/**
 * rmdir - remove directory
 * @param args - the arguments to the command
 */
func rmdir(args []string) {
	if len(args) != 1 {
		fmt.Println("usage: rmdir <dir>")
		return
	}
	dir := args[0]
	err := os.Remove(dir)
	if err != nil {
		fmt.Println(err)
	}
}

/**
 * rm - remove file
 * @param args - the arguments to the command
 */
func rm(args []string) {
	if len(args) != 1 {
		fmt.Println("usage: rm <file>")
		return
	}
	file := args[0]
	err := os.Remove(file)
	if err != nil {
		fmt.Println(err)
	}
}

/**
 * touch - create empty file
 * @param args - the arguments to the command
 */
func touch(args []string) {
	if len(args) != 1 {
		fmt.Println("usage: touch <file>")
		return
	}
	file := args[0]
	f, err := os.Create(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Close()
}

/**
 * cat - concatenate files and print on the standard output
 * @param args - the arguments to the command
 */
func cat(args []string) {
	if len(args) != 1 {
		fmt.Println("usage: cat <file>")
		return
	}
	file := args[0]
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	b := make([]byte, 1024)
	for {
		n, err := f.Read(b)
		if err != nil {
			break
		}
		os.Stdout.Write(b[:n])
	}
	f.Close()
}

/**
 * cp - copy file
 * @param args - the arguments to the command
 */
func cp(args []string) {
	if len(args) != 2 {
		fmt.Println("usage: cp <source> <destination>")
		return
	}
	source := args[0]
	destination := args[1]
	/* usirng a buffer to read the source file */
	buf := make([]byte, 1024)
	sourceFile, err := os.Open(source)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sourceFile.Close()
	destinationFile, err := os.Create(destination)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer destinationFile.Close()
	for {
		n, err := sourceFile.Read(buf)
		if err != nil {
			break
		}
		destinationFile.Write(buf[:n])
	}
}

/**
 * mv - move file
 * @param args - the arguments to the command
 */
func mv(args []string) {
	if len(args) != 2 {
		fmt.Println("usage: mv <source> <destination>")
		return
	}
	source := args[0]
	destination := args[1]
	err := os.Rename(source, destination)
	if err != nil {
		fmt.Println(err)
	}
}

/**
 * exit - exit the shell
 * @param args - the arguments to the command
 */
func exit(args []string) {
	os.Exit(0)
}