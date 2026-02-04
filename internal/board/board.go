// Package board handles mutations
package board

import (
	"slices"

	"github.com/robinwhg/taskmd/internal/markdown"
)

func ToggleTask(task *markdown.Task) {
	task.Checked = !task.Checked
}

func RenameTask(task *markdown.Task, name string) {
	task.Name = name
}

func MoveTaskInColumn(column *markdown.Column, fromIndex, toIndex int) {
	task := column.Tasks[fromIndex]
	column.Tasks = slices.Delete(column.Tasks, fromIndex, fromIndex+1)
	column.Tasks = slices.Insert(column.Tasks, toIndex, task)
}
