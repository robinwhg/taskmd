package session_test

import (
	"io"
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
		FileName: "tasks.md",
		Lines:    []string{"foo", "## To Do"},
		Board:    markdown.Board{Columns: []markdown.Column{{Name: "To Do", Line: 1}}},
	}

	assertError(t, err, nil)

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("got %v, expected %v", got, &want)
	}
}

func TestWriteLines(t *testing.T) {
	fs := fstest.MapFS{
		"tasks.md": {Data: []byte("foo")},
	}

	changes := map[int]session.LineChange{
		0: {Content: "bar"},
	}

	session.WriteLines(fs, "tasks.md", changes)

	file, err := fs.Open("tasks.md")
	assertError(t, err, nil)
	defer func() {
		if closeErr := file.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	fileData, err := io.ReadAll(file)
	assertError(t, err, nil)

	got := string(fileData)
	want := "bar"

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
