// Package session handles the lines and orchestrates other package calls
package session

import (
	"io"
	"strings"

	"github.com/robinwhg/taskmd/internal/board"
	"github.com/robinwhg/taskmd/internal/markdown"
)

const ErrNilWriter = SessionError("write cannot be nil")
	ErrNilWriter = SessionError("writer cannot be nil")

type SessionError string

func (e SessionError) Error() string {
	return string(e)
}

type Session struct {
	Lines []string
	Board markdown.Board
	// TODO: Save writer here
}

func (s *Session) write(writer io.Writer) error {
	if writer == nil {
		return ErrNilWriter
	}

	content := strings.Join(s.Lines, "\n")

	_, err := writer.Write([]byte(content))
	if err != nil {
		return err
	}

	return nil
}

func (s *Session) ToggleTask(columnIndex, taskIndex int, writer io.Writer) error {
	// TODO: Check bounds
	task := &s.Board.Columns[columnIndex].Tasks[taskIndex]

	err := board.ToggleTask(task)
	if err != nil {
		return err
	}

	s.Lines[task.Line] = markdown.RenderTask(*task)

	err = s.write(writer)
	if err != nil {
		return err
	}

	return nil
}

func NewSession(reader io.Reader) (*Session, error) {
	readerData, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(readerData), "\n")

	board, err := markdown.Parse(strings.NewReader(string(readerData)))
	if err != nil {
		return nil, err
	}

	session := Session{Lines: lines, Board: *board}
	return &session, nil
}
