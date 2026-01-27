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
	board := Board{}

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		text := scanner.Text()

		column, found := parseColumn(text)
		if found {
			board.Columns = append(board.Columns, column)
		}
	}

	return &board, nil
}

func parseColumn(text string) (column Column, found bool) {
	after, found := strings.CutPrefix(text, columnPrefix)

	if !found {
		return Column{}, false
	}

	return Column{Name: after}, true
}
