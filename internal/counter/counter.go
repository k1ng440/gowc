package counter

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

// Merge merges the other counter into this one
func (c *Counter) Merge(other *Counter) {
	c.Bytes += other.Bytes
	c.Chars += other.Chars
	c.Lines += other.Lines
	c.MaxLineLength += other.MaxLineLength
	c.Words += other.Words
}
