// This file is part of macfanctl.
//
// Copyright (C) 2013 Roberto Teixeira <robteix@robteix.com>
//
// macfanctl is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// macfanctl is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with macfanctl.  If not, see <http://www.gnu.org/licenses/>.
//

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type Command struct {
	Run func(cmd *Command, args []string)
	UsageLine string
	Help string
	Flag flag.FlagSet
}

func (c *Command) Name() string {
	name := c.UsageLine
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

func (c *Command) Usage() {
	fmt.Fprintf(os.Stderr, "usage: %s\n\n", c.UsageLine)
	os.Exit(2)
}

func (c *Command) Runnable() bool {
	return c.Run != nil
}

var commands = []*Command{
	cmdGet,
	cmdSet,
}

func usage(dst io.Writer) {
	fmt.Fprintf(dst, `
Usage:

	macfanctl command [parameters]

Available commands:
`)
	for _, cmd := range commands {
		fmt.Fprintf(dst, "\t%-15s %s\n", cmd.Name(), cmd.Help)
	}
	fmt.Fprintf(dst, "\nFor more info on each command, use 'macfanctl help command'.\n\n")
	os.Exit(2)
}

func help(args []string) {
	if len(args) == 0 {
		usage(os.Stdout)
		return
	}
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "usage: macfanctl help command\n\nToo many arguments given.\n")
		os.Exit(2)
	}
	arg := args[0]
	for _, cmd := range commands {
		if cmd.Name() == arg {
			fmt.Fprintf(os.Stdout, "\n%s\n\nUsage:\n\tmacfanctl %s\n\n", cmd.Help, cmd.UsageLine)
			return
		}
	}
	fmt.Fprintf(os.Stderr, "Unknown help topic %#q. Run 'macfanctl help'.\n", arg)
	os.Exit(2)
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		usage(os.Stderr)
		return
	}

	if args[0] == "help" {
		help(args[1:])
		return
	}

	for _, cmd := range commands {
		if cmd.Name() == args[0] && cmd.Run != nil {
			cmd.Flag.Usage = func() {cmd.Usage()}
			cmd.Flag.Parse(args[1:])
			args = cmd.Flag.Args()
			cmd.Run(cmd, args)
			return
		}
	}

	fmt.Fprintf(os.Stderr, "Unknown command %#q\n\n", args[0])
	usage(os.Stderr)
}
