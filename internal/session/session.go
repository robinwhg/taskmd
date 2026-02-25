// Package session handles the lines and orchestrates other package calls
package session

import (
	"io"
	"strings"

	"github.com/robinwhg/taskmd/internal/board"
	"github.com/robinwhg/taskmd/internal/markdown"
)

const (
	ErrNilWriter        = SessionError("writer cannot be nil")
	ErrIndexOutOfBounds = SessionError("index out of bounds")
)

type SessionError string

func (e SessionError) Error() string {
	return string(e)
}

type Session struct {
	Lines  []string
	Board  markdown.Board
	writer io.Writer
}

func (s *Session) write() error {
	if s.writer == nil {
		return ErrNilWriter
	}

	content := strings.Join(s.Lines, "\n")

	_, err := s.writer.Write([]byte(content))
	if err != nil {
		return err
	}

	return nil
}

func (s *Session) getColumn(columnIndex int) (*markdown.Column, error) {
	if columnIndex < 0 || columnIndex >= len(s.Board.Columns) {
		return nil, ErrIndexOutOfBounds
	}

	return &s.Board.Columns[columnIndex], nil
}

func (s *Session) getTask(columnIndex, taskIndex int) (*markdown.Task, error) {
	column, err := s.getColumn(columnIndex)
	if err != nil {
		return nil, err
	}

	if taskIndex < 0 || taskIndex >= len(column.Tasks) {
		return nil, ErrIndexOutOfBounds
	}

	return &column.Tasks[taskIndex], nil
}

func (s *Session) ToggleTask(columnIndex, taskIndex int) error {
	task, err := s.getTask(columnIndex, taskIndex)
	if err != nil {
		return err
	}

	if err := board.ToggleTask(task); err != nil {
		return err
	}

	s.Lines[task.Line] = markdown.RenderTask(*task)
	return s.write()
}

func (s *Session) RenameTask(columnIndex, taskIndex int, name string) error {
	task, err := s.getTask(columnIndex, taskIndex)
	if err != nil {
		return err
	}

	if err := board.RenameTask(task, name); err != nil {
		return err
	}

	s.Lines[task.Line] = markdown.RenderTask(*task)
	return s.write()

	// if s.writer == nil {
	// 	return ErrNilWriter
	// }

	// content := strings.Join(s.Lines, "\n")

	// _, err = s.writer.Write([]byte(task.Name))
	// if err != nil {
	// 	return err
	// }
	//
	// return nil
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

	session := Session{Lines: lines, Board: *board, writer: writer}
	return &session, nil
}
