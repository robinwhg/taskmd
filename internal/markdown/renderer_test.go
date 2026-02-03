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

func TestRenderTask(t *testing.T) {
	t.Run("Render unchecked task", func(t *testing.T) {
		input := markdown.Task{Name: "Task A", Checked: false}
		got := markdown.RenderTask(input)
		want := "- [ ] Task A"

		if got != want {
			t.Errorf("got %q, expected %q", got, want)
		}
	})

	t.Run("Render checked task", func(t *testing.T) {
		input := markdown.Task{Name: "Task B", Checked: true}
		got := markdown.RenderTask(input)
		want := "- [x] Task B"

		if got != want {
			t.Errorf("got %q, expected %q", got, want)
		}
	})
}
