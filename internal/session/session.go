// Package session does file operations and orchestrates other package calls
package session

import (
	"io"
	"io/fs"
	"strings"

	"github.com/robinwhg/taskmd/internal/markdown"
)

type Session struct {
	Lines []string
	Board markdown.Board
}

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

	return newSession(file)
}

func newSession(sessionFile io.Reader) (*Session, error) {
	sessionData, err := io.ReadAll(sessionFile)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(sessionData), "\n")
	session := Session{Lines: lines}
	return &session, nil
}
