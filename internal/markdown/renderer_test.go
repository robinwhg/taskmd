package markdown_test

import (
	"testing"

	"github.com/robinwhg/taskmd/internal/markdown"
)

func TestRenderColumn(t *testing.T) {
	input := markdown.Column{Name: "To Do"}
	got := markdown.RenderColumn(input)
	want := "## To Do"

	if got != want {
		t.Errorf("got %q, expected %q", got, want)
	}
}
