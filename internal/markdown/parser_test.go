package markdown_test

import (
	"strings"
	"testing"

	markdown "github.com/robinwhg/taskmd/internal/markdown"
)

const input = `
## To Do

- [ ] Task A

## In Progress

## In Review

## Done
`

func TestReadInput(t *testing.T) {
	board, err := markdown.Parse(strings.NewReader(input))

	assertNoError(t, err)

	if board == nil {
		t.Fatal("Expected non-nil value for Board")
	}
}

func TestParseColumns(t *testing.T) {
	board, err := markdown.Parse(strings.NewReader(input))

	got := board.Columns
	want := []markdown.Column{
		{Name: "To Do"},
		{Name: "In Progress"},
		{Name: "In Review"},
		{Name: "Done"},
	}

	assertNoError(t, err)

	gotLength := len(board.Columns)
	wantLength := len(want)

	if gotLength != wantLength {
		t.Fatalf("got %v columns, expected %v", gotLength, wantLength)
	}

	for index, column := range got {
		gotName := column.Name
		wantName := want[index].Name

		if gotName != wantName {
			t.Fatalf("got column name %v, expected %v", gotName, wantName)
		}
	}
}

func TestParseTasksUnderColumn(t *testing.T) {
	board, err := markdown.Parse(strings.NewReader(input))
	assertNoError(t, err)

	got := board.Columns[0].Tasks
	want := []string{""}

	if len(got) != len(want) {
		t.Fatalf("got %v tasks, expected %v", len(got), len(want))
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error, but didn't want one")
	}
}
