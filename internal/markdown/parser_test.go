package markdown_test

import (
	"strings"
	"testing"

	markdown "github.com/robinwhg/taskmd/internal/markdown"
)

const input = `Hello world`

func TestReadInput(t *testing.T) {
	board, err := markdown.Parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if board == nil {
		t.Fatal("Expected non-nil value for Board")
	}
}
