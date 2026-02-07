// Package board handles mutations
package board

import (
	"slices"

	"github.com/robinwhg/taskmd/internal/markdown"
)

const (
	ErrNilTask       = BoardErr("task cannot be nil")
	ErrEmptyTaskName = BoardErr("task name cannot be empty")
)

type BoardErr string

func (e BoardErr) Error() string {
	return string(e)
}

func ToggleTask(task *markdown.Task) error {
	if task == nil {
		return ErrNilTask
	}

	task.Checked = !task.Checked

	return nil
}

func RenameTask(task *markdown.Task, name string) error {
	if task == nil {
		return ErrNilTask
	}
	if name == "" {
		return ErrEmptyTaskName
	}

	task.Name = name

	return nil
}

func MoveTaskInColumn(column *markdown.Column, fromIndex, toIndex int) {
	if fromIndex == toIndex {
		return
	}

	task := column.Tasks[fromIndex]
	column.Tasks = slices.Delete(column.Tasks, fromIndex, fromIndex+1)
	column.Tasks = slices.Insert(column.Tasks, toIndex, task)
}
