// Package markdown handles the parsing of markdown content
package markdown

import (
	"io"
)

type Board struct{}

func Parse(input io.Reader) (*Board, error) {
	board := Board{}

	return &board, nil
}
