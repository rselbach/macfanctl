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

var cmdGet = &Command{
	Run: runGet,
	UsageLine: "get",
	Help: "get current fan speed",
}

const sysDir = "/sys/devices/platform/applesmc.768"

func runGet(cmd *Command, args []string) {
	fanInput := fmt.Sprintf("%s/fan1_input", sysDir)
	file, err := os.Open(fanInput) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var speed int
	_, err = fmt.Fscanf(file, "%d\n", &speed)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read from %s: %s\n", fanInput, err)
		os.Exit(1)
	}
	fmt.Printf("%d\n", speed)
}
