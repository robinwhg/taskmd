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
	input := "foo\n## To Do"
	output := bytes.Buffer{}
	got, err := session.NewSession(strings.NewReader(input), &output)
	assertError(t, err, nil)

	want := session.Session{
		Lines:  []string{"foo", "## To Do"},
		Board:  markdown.Board{Columns: []markdown.Column{{Name: "To Do", Line: 1}}},
		Writer: &bytes.Buffer{},
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("got %v, expected %v", got, &want)
	}
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

		s, _ := session.NewSession(strings.NewReader(input), &output)

		err := s.ToggleTask(0, 0)
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

		s, _ := session.NewSession(strings.NewReader(input), &output)

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
