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
	"fmt"
	"log"
	"os"
)

var cmdSet = &Command{
	Run: runSet,
	UsageLine: "set [normal|medium|max]",
	Help: "Sets minimal fan speed",
}

type Speed struct {
	Name string
	Speed int
}

var speeds = []*Speed{
	&Speed{"normal", 2000},
	&Speed{"medium", 3500},
	&Speed{"max", 6000},
}

func runSet(cmd *Command, args []string) {
	if len(args) != 1 {
		cmd.Usage()
		os.Exit(-1)
	}
	param := args[0]
	for _, speed := range speeds {
		if param == speed.Name {
			setSpeed(speed.Speed)
			return
		}
	}
	fmt.Fprintf(os.Stderr, "invalid speed specifier: %s\n", param)
	cmd.Usage()
}

func setSpeed(speed int) {
	fanInput := fmt.Sprintf("%s/fan1_min", sysDir)
	file, err := os.OpenFile(fanInput, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Fprintf(file, "%d\n", speed)
}
