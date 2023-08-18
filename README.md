# gowc - A Go Implementation of the Unix Command `wc`

gowc is a command-line tool that serves as a clone of the classic Unix command
`wc`, originally written by Paul Rubin and David MacKenzie. Just like its predecessor,
gowc is designed to provide information about the word, line, character, and byte
count in a given file or input stream.

## Installation

To install gowc, you need to have Go (Golang) installed on your system. If you
don't have it yet, you can download and install it from the official [Go website](https://golang.org/).

Once you have Go installed, you can install gowc using the following steps:

1. Open your terminal.
2. Run the following command to fetch and install gowc:

   ```bash
   go install github.com/k1ng440/gowc/cmd/gowc@latest
   ```
3. After successful installation, you can use gowc from the command line.

## Usage

gowc is used to display the counts of lines, words, characters, and bytes in a file or input stream. The basic syntax is as follows:

```bash
gowc [options ...] [file ...]
```
Here are the available options:

- `-l`: Display the line count.
- `-w`: Display the word count.
- `-c`: Display the character count.
- `-b`: Display the byte count.
- `-m`: Display the character count.
- `--help`: Display help and usage information.
- `--version`: Display current installed version.

You can specify one or more files as arguments. If no file is provided, gowc will read from standard input.

## Examples

1. Display line, word, and character counts for a file:

   ```bash
   gowc file.txt
   ```

2. Display only the word count for multiple files:

   ```bash
   gowc -w file1.txt file2.txt
   ```

3. Count bytes from standard input:

   ```bash
   echo "Hello, gowc!" | gowc -b
   ```

## Contributing

Contributions to gowc are welcome! If you find a bug or have an idea for an improvement,
feel free to open an issue or submit a pull request on the [GitHub repository](https://github.com/k1ng440/gowc).

## License

This project is licensed under the GPL3 License - see the [LICENSE](LICENSE) file for details.

