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
empty slice
single task in slice
check move up and down
*/

func TestMoveTaskInColumn(t *testing.T) {
	tests := []struct {
		name      string
		column    markdown.Column
		fromIndex int
		toIndex   int
		want      markdown.Column
		err       error
	}{
		{
			"Move task up in column",
			makeCol("Task A", "Task B"),
			0,
			1,
			makeCol("Task B", "Task A"),
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
			board.ErrNilTask,
		},
		{
			"toIndex out of lower bounds",
			makeCol("Task A"),
			0,
			-1,
			makeCol("Task A"),
			board.ErrNilTask,
		},
		{
			"fromIndex out of upper bounds",
			makeCol("Task A"),
			2,
			0,
			makeCol("Task A"),
			board.ErrNilTask,
		},
		{
			"toIndex out of upper bounds",
			makeCol("Task A"),
			0,
			2,
			makeCol("Task A"),
			board.ErrNilTask,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := board.MoveTaskInColumn(&tt.column, tt.fromIndex, tt.toIndex)
			assertError(t, err, tt.err)

			if !reflect.DeepEqual(tt.column.Tasks, tt.want.Tasks) {
				t.Errorf("got %v, expected %v", tt.column.Tasks, tt.want.Tasks)
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
