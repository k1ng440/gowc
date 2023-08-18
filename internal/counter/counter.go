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
package counter

import "fmt"

// Counter is a struct that holds the counts for a file
// or a group of files (e.g. the total counts)
type Counter struct {
	Bytes         int
	Chars         int
	Lines         int
	MaxLineLength int
	Words         int
}

// New creates a new counter
func New() *Counter {
	return &Counter{}
}

// AddBytes adds n bytes to the counter
func (c *Counter) AddBytes(n int) {
	c.Bytes += n
}

// AddChars adds n chars to the counter
func (c *Counter) AddChars(n int) {
	c.Chars += n
}

// AddLines adds n lines to the counter
func (c *Counter) AddLines(n int) {
	c.Lines += n
}

// AddMaxLineLength adds n to the max line length
func (c *Counter) AddMaxLineLength(n int) {
	if n > c.MaxLineLength {
		c.MaxLineLength = n
	}
}

// AddWords adds n words to the counter
func (c *Counter) AddWords(n int) {
	c.Words += n
}

// Merge merges the other counters into this one
func (c *Counter) Merge(other ...*Counter) {
	for _, o := range other {
		c.Bytes += o.Bytes
		c.Chars += o.Chars
		c.Lines += o.Lines
		c.MaxLineLength += o.MaxLineLength
		c.Words += o.Words
	}
}

func (c *Counter) String() string {
	return fmt.Sprintf("%d %d %d %d %d", c.Lines, c.Words, c.Bytes, c.Chars, c.MaxLineLength)
}
