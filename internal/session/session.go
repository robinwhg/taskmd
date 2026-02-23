// Package session handles the lines and orchestrates other package calls
package session

import (
	"io"
	"strings"

	"github.com/robinwhg/taskmd/internal/board"
	"github.com/robinwhg/taskmd/internal/markdown"
)

const (
	ErrNilWriter = SessionError("writer cannot be nil")
)

type SessionError string

func (e SessionError) Error() string {
	return string(e)
}

type Session struct {
	Lines  []string
	Board  markdown.Board
	Writer io.Writer
}

func (s *Session) write() error {
	content := strings.Join(s.Lines, "\n")

	_, err := s.Writer.Write([]byte(content))
	if err != nil {
		return err
	}

	return nil
}

func (s *Session) ToggleTask(columnIndex, taskIndex int) error {
	// TODO: Check bounds
	task := &s.Board.Columns[columnIndex].Tasks[taskIndex]

	err := board.ToggleTask(task)
	if err != nil {
		return err
	}

	s.Lines[task.Line] = markdown.RenderTask(*task)

	err = s.write()
	if err != nil {
		return err
	}

	return nil
}

func NewSession(reader io.Reader, writer io.Writer) (*Session, error) {
	readerData, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	if writer == nil {
		return nil, ErrNilWriter
	}

	lines := strings.Split(string(readerData), "\n")

	board, err := markdown.Parse(strings.NewReader(string(readerData)))
	if err != nil {
		return nil, err
	}

	session := Session{Lines: lines, Board: *board, Writer: writer}
	return &session, nil
}
