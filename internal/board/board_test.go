package board_test

import (
	"reflect"
	"testing"

	"github.com/robinwhg/taskmd/internal/board"
	"github.com/robinwhg/taskmd/internal/markdown"
)

func TestToggleTask(t *testing.T) {
	got := markdown.Task{Checked: false}
	want := true
	err := board.ToggleTask(&got)

	assertError(t, err, nil)

	if got.Checked != want {
		t.Errorf("got %v, expected %v", got.Checked, want)
	}
}

func TestRenameTask(t *testing.T) {
	t.Run("Rename a task", func(t *testing.T) {
		got := markdown.Task{Name: "Task A"}
		want := "Task B"
		err := board.RenameTask(&got, want)

		assertError(t, err, nil)

		if got.Name != want {
			t.Errorf("got %v, expected %v", got.Name, want)
		}
	})

	t.Run("Rename a task to an empty string", func(t *testing.T) {
		want := markdown.Task{Name: "Task A"}
		err := board.RenameTask(&want, "")

		assertError(t, err, board.ErrEmptyTaskName)

		if want.Name != "Task A" {
			t.Errorf("got %q, but expected no name change", want.Name)
		}
	})
}

func TestMoveTaskInColumn(t *testing.T) {
	tests := []struct {
		name      string
		got       markdown.Column
		fromIndex int
		toIndex   int
		want      markdown.Column
		err       error
	}{
		{
			"Move task up in column",
			makeCol("Task A", "Task B", "Task C"),
			1,
			0,
			makeCol("Task B", "Task A", "Task C"),
			nil,
		},
		{
			"Move task down in column",
			makeCol("Task A", "Task B", "Task C"),
			1,
			2,
			makeCol("Task A", "Task C", "Task B"),
			nil,
		},
		{
			"Move task to its current index",
			makeCol("Task A"),
			0,
			0,
			makeCol("Task A"),
			nil,
		},
		{
			"fromIndex out of lower bounds",
			makeCol("Task A"),
			-1,
			0,
			makeCol("Task A"),
			board.ErrIndexOutOfBounds,
		},
		{
			"toIndex out of lower bounds",
			makeCol("Task A"),
			0,
			-1,
			makeCol("Task A"),
			board.ErrIndexOutOfBounds,
		},
		{
			"fromIndex out of upper bounds",
			makeCol("Task A"),
			2,
			0,
			makeCol("Task A"),
			board.ErrIndexOutOfBounds,
		},
		{
			"toIndex out of upper bounds",
			makeCol("Task A"),
			0,
			2,
			makeCol("Task A"),
			board.ErrIndexOutOfBounds,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := board.MoveTaskInColumn(&tt.got, tt.fromIndex, tt.toIndex)
			assertError(t, err, tt.err)

			if !reflect.DeepEqual(tt.got.Tasks, tt.want.Tasks) {
				t.Errorf("got %v, expected %v", tt.got.Tasks, tt.want.Tasks)
			}
		})
	}
}

func makeCol(names ...string) markdown.Column {
	tasks := make([]markdown.Task, len(names))
	for i, name := range names {
		tasks[i] = markdown.Task{Name: name}
	}
	return markdown.Column{Tasks: tasks}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
