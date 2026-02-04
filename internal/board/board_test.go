package board_test

import (
	"reflect"
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

func TestRenameTask(t *testing.T) {
	got := markdown.Task{Name: "Task A"}
	want := "Task B"
	board.RenameTask(&got, want)

	if got.Name != want {
		t.Errorf("got %v, expected %v", got.Name, want)
	}
}

func TestMoveTaskInColumn(t *testing.T) {
	got := markdown.Column{Tasks: []markdown.Task{
		{Name: "Task A"},
		{Name: "Task B"},
	}}

	want := markdown.Column{Tasks: []markdown.Task{
		{Name: "Task B"},
		{Name: "Task A"},
	}}

	board.MoveTaskInColumn(&got, 0, 1)

	if !reflect.DeepEqual(got.Tasks, want.Tasks) {
		t.Errorf("got %v, expected %v", got.Tasks, want.Tasks)
	}
}
