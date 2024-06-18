package exercises

import "io"

type Parser struct {
	str        string
	currentRow int64
}

func NewPaser(s string) *Parser {
	return &Parser{str: s}
}

func (parse *Parser) Read(p []byte) (n int, err error) {
	if parse.currentRow >= int64(len(parse.str)) {
		return 0, io.EOF
	}

	n = copy(p, parse.str[parse.currentRow:])
	parse.currentRow += int64(n)
	return
}
