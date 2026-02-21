package session_test

import (
	"bytes"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/robinwhg/taskmd/internal/markdown"
	"github.com/robinwhg/taskmd/internal/session"
)

func TestNewSession(t *testing.T) {
	fs := fstest.MapFS{
		"tasks.md": {Data: []byte("foo\n## To Do")},
	}

	got, err := session.NewSessionFromFS(fs, "tasks.md")
	want := session.Session{
		Lines: []string{"foo", "## To Do"},
		Board: markdown.Board{Columns: []markdown.Column{{Name: "To Do", Line: 1}}},
	}

	assertError(t, err, nil)

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("got %v, expected %v", got, &want)
	}
}

func TestWrite(t *testing.T) {
	input := session.Session{Lines: []string{"foo"}}
	input.Lines[0] = "bar"
	buffer := bytes.Buffer{}
	err := input.Write(&buffer)
	assertError(t, err, nil)
	got := buffer.String()
	want := "bar"

	if got != want {
		t.Errorf("got %q, expected %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
