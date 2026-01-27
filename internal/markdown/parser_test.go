package markdown_test

import (
	"strings"
	"testing"

	markdown "github.com/robinwhg/taskmd/internal/markdown"
)

const input = `
## To Do

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
		t.Fatalf("expected %v columns, found %v", gotLength, wantLength)
	}

	for index, column := range got {
		gotName := column.Name
		wantName := want[index].Name

		if gotName != wantName {
			t.Fatalf("expected column name %v, got %v", gotName, wantName)
		}
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error, but didn't want one")
	}
}
