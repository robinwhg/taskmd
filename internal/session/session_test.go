package session_test

import (
	"bytes"
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/robinwhg/taskmd/internal/markdown"
	"github.com/robinwhg/taskmd/internal/session"
)

func TestNewSession(t *testing.T) {
	t.Run("Test new session", func(t *testing.T) {
		input := "foo\n## To Do"
		output := bytes.Buffer{}
		got, err := session.NewSession(strings.NewReader(input), &output)
		assertError(t, err, nil)
		if got == nil {
			t.Fatal("expected non-nil session")
		}

		wantLines := []string{"foo", "## To Do"}
		if !reflect.DeepEqual(got.Lines, wantLines) {
			t.Errorf("got %v, expected %v", got.Lines, wantLines)
		}

		wantBoard := markdown.Board{Columns: []markdown.Column{{Name: "To Do", Line: 1}}}
		if !reflect.DeepEqual(got.Board, wantBoard) {
			t.Errorf("got %v, expected %v", got.Board, wantBoard)
		}

		if got.Writer == nil {
			t.Errorf("expected non-nil writer")
		}
	})

	t.Run("Test nil writer", func(t *testing.T) {
		input := "foo\n## To Do"
		_, err := session.NewSession(strings.NewReader(input), nil)
		assertError(t, err, session.ErrNilWriter)
	})
}

var errWrite = errors.New("write failed")

type stubFailingWriter struct{}

func (s stubFailingWriter) Write(p []byte) (n int, err error) {
	return 0, errWrite
}

func TestToggleTask(t *testing.T) {
	t.Run("Toggle a task", func(t *testing.T) {
		input := "- [ ] Task A"
		output := bytes.Buffer{}

		s, err := session.NewSession(strings.NewReader(input), &output)
		assertError(t, err, nil)

		err = s.ToggleTask(0, 0)
		assertError(t, err, nil)

		got := output.String()
		want := "- [x] Task A"

		if got != want {
			t.Errorf("got %v, expected %v", output, want)
		}
	})

	t.Run("Failing Write", func(t *testing.T) {
		input := "- [ ] Task A"
		output := stubFailingWriter{}

		s, err := session.NewSession(strings.NewReader(input), &output)
		assertError(t, err, nil)

		err = s.ToggleTask(0, 0)
		assertError(t, err, errWrite)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
