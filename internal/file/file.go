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

package file

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"unicode/utf8"

	"github.com/k1ng440/gowc/internal/counter"
)

var (
	ErrIsDir      = errors.New("is a directory")
	ErrNotFound   = errors.New("no such file")
	ErrCannotOpen = errors.New("cannot open file")
)

type File struct {
	reader *os.File
	Name   string
}

func New(file string) (*File, error) {
	f := &File{
		Name: file,
	}

	// check if the file exists
	if err := f.Validate(); err != nil {
		return nil, err
	}

	return f, nil
}

func (f *File) Count() (*counter.Counter, error) {
	counter := counter.New()

	if f.reader == nil {
		if err := f.Open(); err != nil {
			return nil, err
		}

		defer func() {
			if err := f.Close(); err != nil {
				log.Println(err)
			}
		}()
	}

	scanner := bufio.NewScanner(f.reader)
	for scanner.Scan() {
		line := scanner.Bytes()
		runeCount := utf8.RuneCount(line)

		counter.AddLines(1)
		counter.AddBytes(len(line) + 1) // +1 for the newline character
		counter.AddChars(runeCount + 1) // +1 for the newline character
		counter.AddWords(len(bytes.Fields(line)))
		counter.AddMaxLineLength(runeCount)
	}

	err := scanner.Err()
	if err != nil && !errors.Is(io.EOF, err) {
		return nil, err
	}

	return counter, nil
}

// Validate checks if the file exists and if it is a file and not a directory
func (f *File) Validate() error {
	info, err := os.Stat(f.Name)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%w: %s", ErrNotFound, f.Name)
		}

		// Unknown error
		return err
	}

	if info.IsDir() {
		return fmt.Errorf("%w: %s", ErrIsDir, f.Name)
	}

	return nil
}

func (f *File) Open() (err error) {
	f.reader, err = os.Open(f.Name)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrCannotOpen, f.Name)
	}

	return nil
}

func (f *File) Close() error {
	return f.reader.Close()
}
