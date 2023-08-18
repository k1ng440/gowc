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
)

func printHelp() {
	w := flag.CommandLine.Output()
	fmt.Fprintln(w, "NAME")
	fmt.Fprintln(w, "        wc - print newline, word, and byte counts for each file")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "SYNOPSIS")
	fmt.Fprintln(w, "        wc [OPTION]... [FILE]...")
	fmt.Fprintln(w, "        wc [OPTION]... --files0-from=F")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "DESCRIPTION")
	fmt.Fprintln(w, "        Print newline, word, and byte counts for each FILE,"+
		"and a total line if more than one FILE is specified.",
	)
	fmt.Fprintln(w, "        A word is a non-zero-length sequence of characters delimited by white space.")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "        With no FILE, or when FILE is -, read standard input.")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "        The options below may be used to select which counts are"+
		"printed, always in the following order: newline, word, character, byte, "+
		"maximum line length.")
	fmt.Fprintln(w)

	flag.CommandLine.PrintDefaults()
	fmt.Fprintln(w)

	fmt.Fprintln(w, "AUTHOR")
	fmt.Fprintln(w, "       Written by Aasaduzzaman Pavel.")
}
