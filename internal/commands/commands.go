/*
*    gowc
*    Copyright (C) 2023 Aasaduzzaman Pavel contact@iampavel.dev
*
*    This program is free software: you can redistribute it and/or modify
*    it under the terms of the GNU General Public License as published by
*    the Free Software Foundation, either version 3 of the License, or
*    (at your option) any later version.
*
*    This program is distributed in the hope that it will be useful,
*    but WITHOUT ANY WARRANTY; without even the implied warranty of
*    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*    GNU General Public License for more details.
*
*    You should have received a copy of the GNU General Public License
*    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package commands

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/k1ng440/gowc/internal/arguments"
	"github.com/k1ng440/gowc/internal/counter"
	"github.com/k1ng440/gowc/internal/file"
	"github.com/k1ng440/gowc/internal/printer"
)

// Version is the version of the program
// This variable is set at compile time
var Version = "0.0.1"

func Run() {
	args := arguments.ParseArguments()

	if args.HasHelp() {
		flag.Usage()
		return
	}

	if args.HasVersion() {
		fmt.Println("gowc " + Version)
		return
	}

	// TODO: Read from stdin
	if len(args.Files) == 0 {
		fmt.Println("TODO: read from stdin")
		os.Exit(69)
	}

	handleFiles(args)
}

func handleFiles(args arguments.Arguments) {
	counters := make([]*counter.Counter, 0, len(args.Files))
	printer := printer.New(args)

	for _, fp := range args.Files {
		f, err := file.New(fp)
		if err != nil {
			if errors.Is(err, file.ErrNotFound) {
				fmt.Println("File does not exist:", fp)
				continue
			}

			if errors.Is(err, file.ErrIsDir) {
				fmt.Println("File is a directory:", fp)
				continue
			}

			fmt.Println(err)
			os.Exit(1)
		}

		counts, err := f.Count()
		if err != nil {
			fmt.Printf("Error: %s, %s\n", fp, err)
			continue
		}

		printer.Print(fp, counts)
		counters = append(counters, counts)
	}

	if len(counters) > 1 {
		total := counter.New()
		total.Merge(counters...)
		printer.Print("total", total)
	}
}
