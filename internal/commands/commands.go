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
	"flag"
	"fmt"
	"os"

	"github.com/k1ng440/gowc/internal/counter"
	"github.com/k1ng440/gowc/internal/file"
)

// Version is the version of the program
// This variable is set at compile time
var Version = "0.0.1"

type Arguments struct {
	Files0From    string
	Files         []string
	Bytes         bool
	Chars         bool
	Lines         bool
	MaxLineLength bool
	Words         bool
	Help          bool
	Version       bool
}

func Run() {
	args := parseArguments()

	if args.HasHelp() {
		flag.Usage()
		return
	}

	if args.HasVersion() {
		fmt.Println("wc " + Version)
		return
	}

	// TODO: Read from stdin
	if len(args.Files) == 0 {
		fmt.Println("TODO: read from stdin")
		os.Exit(69)
	}

	counters := make([]*counter.Counter, 0)

	for _, fp := range args.Files {
		f, err := file.New(fp)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		counts, err := f.Count()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("%#v\n", counts)

		counters = append(counters, counts)
	}

	if len(counters) > 1 {
		total := counter.New()
		for _, c := range counters {
			total.Merge(c)
		}

		fmt.Printf("Total: %#v\n", total)
	}
}

func parseArguments() Arguments {
	args := Arguments{}

	flag.StringVar(&args.Files0From, "files0-from", "", "read input from the "+
		"files specified by NUL-terminated names in file F; If F is - then read"+
		" names from standard input",
	)
	flag.BoolVar(&args.Bytes, "c", false, "print the byte counts")
	flag.BoolVar(&args.Chars, "m", false, "print the character counts")
	flag.BoolVar(&args.Lines, "l", false, "print the newline counts")
	flag.BoolVar(&args.MaxLineLength, "L", false, "print the maximum display width")
	flag.BoolVar(&args.Words, "w", false, "print the word counts")
	flag.BoolVar(&args.Help, "help", false, "display this help and exit")
	flag.BoolVar(&args.Version, "version", false, "output version information and exit")

	flag.Parse()
	args.Files = flag.Args()
	flag.Usage = printHelp
	return args
}

func (a Arguments) HasFiles() bool {
	return len(a.Files) > 0
}

func (a Arguments) HasFlags() bool {
	return a.Bytes || a.Chars || a.Lines || a.MaxLineLength || a.Words
}

func (a Arguments) HasHelp() bool {
	return a.Help
}

func (a Arguments) HasVersion() bool {
	return a.Version
}

func (a Arguments) HasFiles0From() bool {
	return a.Files0From != ""
}

func (a Arguments) HasStdin() bool {
	return len(a.Files) == 0 && !a.HasFiles0From()
}
