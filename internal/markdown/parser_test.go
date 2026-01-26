package markdown_test

import (
	"reflect"
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

	assertNoError(t, err)

	t.Run("Create columns", func(t *testing.T) {
		got := len(board.Columns)

		want := 4

		if got != want {
			t.Fatalf("expected %v columns, found %v", want, got)
		}
	})

	t.Run("Create columns with title", func(t *testing.T) {
		got := board.Columns

		want := []markdown.Column{
			{Name: "To Do"},
			{Name: "In Progress"},
			{Name: "In Review"},
			{Name: "Done"},
		}

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("expected columns %v got %v", want, got)
		}
	})
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error, but didn't want one")
	}
}
