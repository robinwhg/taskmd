package board_test

import (
	"reflect"
	"testing"

	"github.com/robinwhg/taskmd/internal/board"
	"github.com/robinwhg/taskmd/internal/markdown"
)

func TestToggleTask(t *testing.T) {
	got := markdown.Task{Checked: false}
	err := board.ToggleTask(&got)
	want := true

	assertNoError(t, err)

	if got.Checked != want {
		t.Errorf("got %v, expected %v", got.Checked, want)
	}
}

func TestRenameTask(t *testing.T) {
	t.Run("Rename a task", func(t *testing.T) {
		got := markdown.Task{Name: "Task A"}
		want := "Task B"
		err := board.RenameTask(&got, want)

		assertNoError(t, err)

		if got.Name != want {
			t.Errorf("got %v, expected %v", got.Name, want)
		}
	})

	t.Run("Rename a task to an empty string", func(t *testing.T) {
		input := markdown.Task{Name: "Task A"}
		got := board.RenameTask(&input, "")

		assertError(t, got, board.ErrEmptyTaskName)

		if input.Name != "Task A" {
			t.Errorf("got %q, but expected no name change", input.Name)
		}
	})
}

/* TODO: Test
fromIndex == sameIndex
index out of bounds
empty slice
single task in slice
*/

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

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
