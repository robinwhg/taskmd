// Package session does file operations and orchestrates other package calls
package session

import (
	"io"
	"io/fs"
	"strings"

	"github.com/robinwhg/taskmd/internal/markdown"
)

type Session struct {
	FileName string
	Lines    []string
	Board    markdown.Board
}

func (s *Session) Write(writer io.Writer) error {
	content := strings.Join(s.Lines, "\n")

	_, err := writer.Write([]byte(content))
	if err != nil {
		return err
	}

	return nil
}

// Why even pass anything if I have the filename?
// NOTE: WriteLines should be private, but need to test it somehow

func NewSessionFromFS(fileSystem fs.FS, fileName string) (session *Session, err error) {
	file, err := fileSystem.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	return newSession(file, fileName)
}

func newSession(sessionFile io.Reader, fileName string) (*Session, error) {
	sessionData, err := io.ReadAll(sessionFile)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(sessionData), "\n")

	board, err := markdown.Parse(strings.NewReader(string(sessionData)))
	if err != nil {
		return nil, err
	}

	session := Session{FileName: fileName, Lines: lines, Board: *board}
	return &session, nil
}
