package markdown_test

import (
	"strings"
	"testing"

	markdown "github.com/robinwhg/taskmd/internal/markdown"
)

const input = `
## To Do

- [ ] Task A
- [x] Task B

## In Progress

## In Review

## Done

`

func TestParseReturnsBoard(t *testing.T) {
	board, err := markdown.Parse(strings.NewReader(input))

	assertNoError(t, err)

	if board == nil {
		t.Fatal("Expected non-nil value for Board")
	}
}

func TestParseColumns(t *testing.T) {
	board, err := markdown.Parse(strings.NewReader(input))
	assertNoError(t, err)

	got := board.Columns
	want := []string{"To Do", "In Progress", "In Review", "Done"}

	if len(got) != len(want) {
		t.Fatalf("got %v columns, expected %v", len(got), len(want))
	}

	for i, column := range got {
		if column.Name != want[i] {
			t.Fatalf("got column name %q, expected %q", column.Name, want[i])
		}
	}
}

func TestParseTasksUnderColumn(t *testing.T) {
	board, err := markdown.Parse(strings.NewReader(input))
	assertNoError(t, err)

	got := board.Columns[0].Tasks
	want := []markdown.Task{{Name: "Task A", Checked: false}, {Name: "Task B", Checked: true}}

	if len(got) != len(want) {
		t.Fatalf("got %v tasks, expected %v", len(got), len(want))
	}

	for i, task := range got {
		if task.Checked != want[i].Checked {
			t.Fatalf("got Checked %v, expected %v", task.Checked, want[i].Checked)
		}

		if task.Name != want[i].Name {
			t.Fatalf("got Name %q, expected %q", task.Name, want[i].Name)
		}
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error, but didn't want one")
	}
}
