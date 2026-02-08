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

func TestInsertTask(t *testing.T) {
	t.Run("Insert task to nil column", func(t *testing.T) {
		err := board.InsertTask(nil, 0, markdown.Task{Name: "Task A"})
		assertError(t, err, board.ErrNilColumn)
	})

	tests := []struct {
		name  string
		got   markdown.Column
		index int
		task  markdown.Task
		want  markdown.Column
		err   error
	}{
		{
			"Insert a task",
			makeCol("Task A"),
			1,
			markdown.Task{Name: "Task B"},
			makeCol("Task A", "Task B"),
			nil,
		},
		{
			"Index out of lower bounds",
			makeCol("Task A"),
			-1,
			markdown.Task{Name: "Task B"},
			makeCol("Task A"),
			board.ErrIndexOutOfBounds,
		},
		{
			"Index out of upper bounds",
			makeCol("Task A"),
			2,
			markdown.Task{Name: "Task B"},
			makeCol("Task A"),
			board.ErrIndexOutOfBounds,
		},
		{
			"Insert task with empty name",
			makeCol("Task A"),
			1,
			markdown.Task{Name: ""},
			makeCol("Task A"),
			board.ErrEmptyTaskName,
		},
	}

	for _, tt := range tests {
		err := board.InsertTask(&tt.got, tt.index, tt.task)
		assertError(t, err, tt.err)

		if !reflect.DeepEqual(tt.got.Tasks, tt.want.Tasks) {
			t.Errorf("got %v, expected %v", tt.got.Tasks, tt.want.Tasks)
		}
	}
}

func TestDeleteTask(t *testing.T) {
	t.Run("Delete task in nil column", func(t *testing.T) {
		err := board.DeleteTask(nil, 0)
		assertError(t, err, board.ErrNilColumn)
	})

	tests := []struct {
		name  string
		got   markdown.Column
		index int
		want  markdown.Column
		err   error
	}{
		{
			"Delete a Task",
			makeCol("Task A"),
			0,
			makeCol(),
			nil,
		},
		{
			"Index out of lower bounds",
			makeCol("Task A"),
			-1,
			makeCol("Task A"),
			board.ErrIndexOutOfBounds,
		},
		{
			"Index out of upper bounds",
			makeCol("Task A"),
			1,
			makeCol("Task A"),
			board.ErrIndexOutOfBounds,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := board.DeleteTask(&tt.got, tt.index)
			assertError(t, err, tt.err)

			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("got %v, expected %v", tt.got, tt.want)
			}
		})
	}
}

func TestMoveTaskInColumn(t *testing.T) {
	t.Run("Move task in a nil column", func(t *testing.T) {
		err := board.MoveTaskInColumn(nil, 0, 1)
		assertError(t, err, board.ErrNilColumn)
	})

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

func TestMoveTaskToColumn(t *testing.T) {
	t.Run("Move task from nil column", func(t *testing.T) {
		err := board.MoveTaskToColumn(nil, 0, &markdown.Column{}, 0)
		assertError(t, err, board.ErrNilColumn)
	})
	t.Run("Move task to nil column", func(t *testing.T) {
		err := board.MoveTaskToColumn(&markdown.Column{}, 0, nil, 0)
		assertError(t, err, board.ErrNilColumn)
	})

	tests := []struct {
		name      string
		got       [2]markdown.Column
		fromIndex int
		toIndex   int
		want      [2]markdown.Column
		err       error
	}{
		{
			"Move task to column",
			[2]markdown.Column{makeCol("Task A"), makeCol()},
			0,
			0,
			[2]markdown.Column{makeCol(), makeCol("Task A")},
			nil,
		},
		{
			"fromIndex out of lower bounds",
			[2]markdown.Column{makeCol("Task A"), makeCol()},
			-1,
			0,
			[2]markdown.Column{makeCol("Task A"), makeCol()},
			board.ErrIndexOutOfBounds,
		},
		{
			"fromIndex out of upper bounds",
			[2]markdown.Column{makeCol("Task A"), makeCol()},
			2,
			0,
			[2]markdown.Column{makeCol("Task A"), makeCol()},
			board.ErrIndexOutOfBounds,
		},
		{
			"toIndex out of lower bounds",
			[2]markdown.Column{makeCol("Task A"), makeCol()},
			0,
			-1,
			[2]markdown.Column{makeCol("Task A"), makeCol()},
			board.ErrIndexOutOfBounds,
		},
		{
			"toIndex out of upper bounds",
			[2]markdown.Column{makeCol("Task A"), makeCol()},
			0,
			2,
			[2]markdown.Column{makeCol("Task A"), makeCol()},
			board.ErrIndexOutOfBounds,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := board.MoveTaskToColumn(&tt.got[0], tt.fromIndex, &tt.got[1], tt.toIndex)
			assertError(t, err, tt.err)

			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("got %v, expected %v", tt.got, tt.want)
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
