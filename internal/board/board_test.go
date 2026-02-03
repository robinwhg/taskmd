package board_test

import (
	"testing"

	"github.com/robinwhg/taskmd/internal/board"
	"github.com/robinwhg/taskmd/internal/markdown"
)

func TestToggleTask(t *testing.T) {
	got := markdown.Task{Checked: false}
	board.ToggleTask(&got)
	want := true

	if got.Checked != want {
		t.Errorf("got %v, expected %v", got.Checked, want)
	}
}
