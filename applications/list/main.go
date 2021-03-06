/*
   Copyright (C) 2014  Oscar Campos <oscar.campos@member.fsf.org>

   This program is free software; you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation; either version 2 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License along
   with this program; if not, write to the Free Software Foundation, Inc.,
   51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.

   See LICENSE file for more details.
*/

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DamnWidget/VenGO/commands"
	"github.com/DamnWidget/VenGO/utils"
	flag "github.com/ogier/pflag"
)

var all, installed, nonInstalled, asJson, help bool

func init() {
	flag.BoolVarP(
		&installed, "installed", "i", true, "Display installed Go versions")
	flag.BoolVarP(
		&all, "all", "a", false, "Dsiplay installed and available Go versions")
	flag.BoolVarP(&asJson, "json", "j", false, "Display results as JSON")
	flag.BoolVarP(
		&nonInstalled, "non-installed", "n", false,
		"Display non installed but available Go versions",
	)
	flag.BoolVarP(&help, "help", "h", false, "display help message")
	flag.Parse()
}

// main function entry point
func main() {
	if help {
		displayHelp()
		os.Exit(1)
	}

	// build the list object based on the given options
	options := buildCommandListOptions()
	nl := commands.NewList(options...)
	data, err := nl.Run()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println(data)
}

// build the command list options based in the passed flags
func buildCommandListOptions() []func(*commands.List) {
	options := []func(l *commands.List){}
	if asJson {
		options = append(options, func(l *commands.List) {
			l.DisplayAs = commands.Json
		})
	}
	if all {
		options = append(options, func(l *commands.List) {
			l.ShowBoth = true
		})
	} else {
		if nonInstalled {
			options = append(options, func(l *commands.List) {
				l.ShowNotInstalled = true
			})
		}
		if installed {
			options = append(options, func(l *commands.List) {
				l.ShowInstalled = true
			})
		} else {
			options = append(options, func(l *commands.List) {
				l.ShowInstalled = false
			})
		}
	}

	return options
}

// display help message
func displayHelp() {
	fmt.Println(fmt.Sprintf(`%s: vengo list [options]
    -a, --all	 			Display installed and available Go versions
    -i, --installed			Display installed go versions
    -n, --non-installed			Display available go versions
    -j, --json				Display the results in JSON format

    -h, --help				Display this message
	`, utils.Ok("Usage")))
}
