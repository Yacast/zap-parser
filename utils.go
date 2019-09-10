package zapparser

import (
	"bytes"
	"os"
)

// FromFile creates a new scanner from a file path.
func FromFile(path string) (*Parser, error) {
	f, err := os.Open(path)
	parser := NewParser(f)
	parser.OnClose(func() {
		f.Close()
	})
	return parser, err
}

// FromFile creates a new scanner from a byte slice.
func FromBytes(bs []byte) *Parser {
	return NewParser(bytes.NewReader(bs))
}

// FromString creates a new scanner from a string.
func FromString(s string) *Parser {
	return FromBytes([]byte(s))
}
