package session_test

import (
	"bytes"
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/robinwhg/taskmd/internal/board"
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
	tests := []struct {
		name        string
		columnIndex int
		taskIndex   int
		input       string
		want        string
		err         error
	}{
		{
			"Toggle an unchecked task",
			0,
			0,
			"- [ ] Task A",
			"- [x] Task A",
			nil,
		}, {
			"Toggle a checked task",
			0,
			0,
			"- [x] Task A",
			"- [ ] Task A",
			nil,
		}, {
			"Column index out of lower bounds",
			-1,
			0,
			"- [x] Task A",
			"",
			session.ErrIndexOutOfBounds,
		}, {
			"Column index out of upper bounds",
			1,
			0,
			"- [x] Task A",
			"",
			session.ErrIndexOutOfBounds,
		}, {
			"Task index out of lower bounds",
			0,
			-1,
			"- [x] Task A",
			"",
			session.ErrIndexOutOfBounds,
		}, {
			"Task index out of upper bounds",
			0,
			1,
			"- [x] Task A",
			"",
			session.ErrIndexOutOfBounds,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := bytes.Buffer{}
			s, err := session.NewSession(strings.NewReader(tt.input), &writer)
			assertError(t, err, nil)

			err = s.ToggleTask(tt.columnIndex, tt.taskIndex)
			assertError(t, err, tt.err)

			got := writer.String()
			if got != tt.want {
				t.Errorf("got %v, expected %v", got, tt.want)
			}
		})
	}

	t.Run("Failing Write", func(t *testing.T) {
		input := "- [ ] Task A"
		output := stubFailingWriter{}

		s, err := session.NewSession(strings.NewReader(input), &output)
		assertError(t, err, nil)

		err = s.ToggleTask(0, 0)
		assertError(t, err, errWrite)
	})
}

func TestRenameTask(t *testing.T) {
	t.Run("Rename a task", func(t *testing.T) {
		writer := bytes.Buffer{}
		s, err := session.NewSession(strings.NewReader("- [ ] Task A"), &writer)
		assertError(t, err, nil)

		err = s.RenameTask(0, 0, "Task B")
		assertError(t, err, nil)

		got := writer.String()
		want := "- [ ] Task B"
		if got != want {
			t.Errorf("got %v, expected %v", got, want)
		}
	})

	t.Run("Rename a task to an empty string", func(t *testing.T) {
		writer := bytes.Buffer{}
		s, err := session.NewSession(strings.NewReader("- [ ] Task A"), &writer)
		assertError(t, err, nil)

		err = s.RenameTask(0, 0, "")
		assertError(t, err, board.ErrEmptyTaskName)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
