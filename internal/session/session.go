// Package session handles the lines and orchestrates other package calls
package session

import (
	"io"
	"strings"

	"github.com/robinwhg/taskmd/internal/markdown"
)

type Session struct {
	Lines []string
	Board markdown.Board
}

func (s *Session) Write(writer io.Writer) error {
	content := strings.Join(s.Lines, "\n")

	_, err := writer.Write([]byte(content))
	if err != nil {
		return err
	}

	return nil
}

func NewSession(sessionFile io.Reader) (*Session, error) {
	sessionData, err := io.ReadAll(sessionFile)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(sessionData), "\n")

	board, err := markdown.Parse(strings.NewReader(string(sessionData)))
	if err != nil {
		return nil, err
	}

	session := Session{Lines: lines, Board: *board}
	return &session, nil
}
