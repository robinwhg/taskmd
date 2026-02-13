package session_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/robinwhg/taskmd/internal/session"
)

func TestNewSession(t *testing.T) {
	fs := fstest.MapFS{
		"tasks.md": {Data: []byte("foo\nbar")},
	}

	got, err := session.NewSessionFromFS(fs, "tasks.md")
	want := session.Session{Lines: []string{"foo", "bar"}}

	assertError(t, err, nil)

	if !reflect.DeepEqual(got, &want) {
		t.Errorf("got %v, expected %v", got, &want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
