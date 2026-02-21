package session_test

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/robinwhg/taskmd/internal/markdown"
	"github.com/robinwhg/taskmd/internal/session"
)

func TestNewSession(t *testing.T) {
	input := "foo\n## To Do"
	got, err := session.NewSession(strings.NewReader(input))
	assertError(t, err, nil)

	want := session.Session{
		Lines: []string{"foo", "## To Do"},
		Board: markdown.Board{Columns: []markdown.Column{{Name: "To Do", Line: 1}}},
	}

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("got %v, expected %v", got, &want)
	}
}

func TestWriteSession(t *testing.T) {
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
