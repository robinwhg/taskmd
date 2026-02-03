// Package board handles mutations
package board

import "github.com/robinwhg/taskmd/internal/markdown"

func ToggleTask(task *markdown.Task) {
	task.Checked = !task.Checked
}
