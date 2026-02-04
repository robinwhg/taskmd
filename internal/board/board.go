// Package board handles mutations
package board

import (
	"errors"
	"slices"

	"github.com/robinwhg/taskmd/internal/markdown"
)

func ToggleTask(task *markdown.Task) error {
	if task == nil {
		return errors.New("task is nil")
	}

	task.Checked = !task.Checked

	return nil
}

func RenameTask(task *markdown.Task, name string) {
	task.Name = name
}

func MoveTaskInColumn(column *markdown.Column, fromIndex, toIndex int) {
	task := column.Tasks[fromIndex]
	column.Tasks = slices.Delete(column.Tasks, fromIndex, fromIndex+1)
	column.Tasks = slices.Insert(column.Tasks, toIndex, task)
}
