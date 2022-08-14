# Unix Command Line Interpreter - GO implementation

## Introduction

This is a simple implementation of the Unix command line interpreter in Go.

## Features

### built-in commands

- `cd`: change directory
- `pwd`: print working directory
- `ls`: list files in current directory
- `cat`: concatenate files and print on standard output
- `rm`: remove files
- `mkdir`: create directories
- `rmdir`: remove directories
- `touch`: create empty files
- `mv`: move files
- `cp`: copy files
- `grep`: search for a pattern in files
- `wc`: count lines, words and bytes in files
- `sed`: search and replace text in files
- `history`: show command history
- `clear`: clear command history
- `exit`: exit the interpreter

### external commands

- `man`: show manual page for a command
- `help`: show help for a command
- `which`: show path for a command
- `whereis`: show path for a command
- `type`: show type of a file
- `locate`: show path for a file
- `apropos`: show all commands matching a pattern
- `whatis`: show description for a command
- `alias`: show or set aliases
- `unalias`: remove aliases
- `set`: show or set options
- `unset`: remove options
- `export`: show or set environment variables
- `read`: read a file into the current shell
- `.`: source a file


## Installation

### go get

```bash
go get -u github.com/Aymen-haddaji-hub/Unix_Go_Shell
```

### go install

```bash
go install github.com/Aymen-haddaji-hub/Unix_Go_Shell
```

## Usage

```bash
$ go run main.go helper.go builtin.go
```

## Build

```bash
$ go build -o go-shell main.go helper.go builtin.go
```

## Run

```bash
$ ./go-shell
```

## License

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

## Authors

- AymenHaddaji < [@AymenHaddaji](aymensystem7@gmail.com) >

## Future improvements

- add support for external commands
- add support for pipes
- add support for redirections
- add support for multiple commands
- add support for history
- add support for aliases
- add support for environment variables
- add support for functions


## Contributors

- [@AymenHaddaji](aymensystem7@gmail.com)


## Bugs

- please report any bugs to [@AymenHaddaji](aymensystem7@gmail.com)
or create an issue on [Github](https://github.com/Aymen-haddaji-hub/Unix_Go_Shell)
