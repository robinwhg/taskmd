package markdown_test

import (
	"strings"
	"testing"

	markdown "github.com/robinwhg/taskmd/internal/markdown"
)

const input = `# My Task Board

## To Do

- [ ] Task A
- [x] Task B

## In Progress

## Done

`

func TestParseReturnsBoard(t *testing.T) {
	board, err := markdown.Parse(strings.NewReader(input))

	assertNoError(t, err)

	if board == nil {
		t.Fatal("Expected non-nil value for Board")
	}
}

func TestParseTitle(t *testing.T) {
	board, err := markdown.Parse(strings.NewReader(input))
	assertNoError(t, err)

	got := board.Title
	want := "My Task Board"

	if got != want {
		t.Errorf("got title %q, expected %q", got, want)
	}
}

func TestParseColumns(t *testing.T) {
	board, err := markdown.Parse(strings.NewReader(input))
	assertNoError(t, err)

	got := board.Columns
	want := []markdown.Column{{Name: "To Do", Line: 2}, {Name: "In Progress", Line: 7}, {Name: "Done", Line: 9}}

	if len(got) != len(want) {
		t.Fatalf("got %v columns, expected %v", len(got), len(want))
	}

	for i, column := range got {
		if column.Name != want[i].Name {
			t.Fatalf("got column name %q, expected %q", column.Name, want[i].Name)
		}
		if column.Line != want[i].Line {
			t.Fatalf("got column on line %v, expected %v", column.Line, want[i].Line)
		}
	}
}

func TestParseTasksUnderColumn(t *testing.T) {
	board, err := markdown.Parse(strings.NewReader(input))
	assertNoError(t, err)

	got := board.Columns[0].Tasks
	want := []markdown.Task{{Name: "Task A", Checked: false, Line: 4}, {Name: "Task B", Checked: true, Line: 5}}

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

		if task.Line != want[i].Line {
			t.Fatalf("got Line %v, expected %v", task.Line, want[i].Line)
		}
	}
}

func TestParseTasksUnderNoColumn(t *testing.T) {
	input := `
- [ ] Orphan Task

## To Do

- [ ] Task A
`
	board, err := markdown.Parse(strings.NewReader(input))
	assertNoError(t, err)

	if board.Columns[0].Name != "Uncategorized" {
		t.Fatalf("got column with name %q, but expected Uncategorized", board.Columns[0].Name)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error, but didn't want one")
	}
}
