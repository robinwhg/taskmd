// Package markdown handles the parsing of markdown content
package markdown

import (
	"bufio"
	"io"
	"strings"
)

const (
	columnPrefix = "## "
)

type Column struct {
	Name string
}

type Board struct {
	Columns []Column
}

func Parse(input io.Reader) (*Board, error) {
	board := Board{Columns: ParseColumns(input)}

	return &board, nil
}

func ParseColumns(input io.Reader) []Column {
	scanner := bufio.NewScanner(input)

	columns := []Column{}

	for scanner.Scan() {
		text := scanner.Text()

		after, found := strings.CutPrefix(text, columnPrefix)

		if found {
			columns = append(columns, Column{Name: after})
		}
	}

	return columns
}
