// Package board handles mutations
package board

import (
	"slices"

	"github.com/robinwhg/taskmd/internal/markdown"
)

const (
	ErrNilColumn        = BoardErr("column cannot be nil")
	ErrNilTask          = BoardErr("task cannot be nil")
	ErrEmptyTaskName    = BoardErr("task name cannot be empty")
	ErrIndexOutOfBounds = BoardErr("index out of bounds")
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

func InsertTask(column *markdown.Column, index int, task markdown.Task) error {
	if column == nil {
		return ErrNilColumn
	}
	if task.Name == "" {
		return ErrEmptyTaskName
	}
	if index < 0 || index > len(column.Tasks) {
		return ErrIndexOutOfBounds
	}

	column.Tasks = slices.Insert(column.Tasks, index, task)

	return nil
}

func MoveTaskInColumn(column *markdown.Column, fromIndex, toIndex int) error {
	if fromIndex == toIndex {
		return nil
	}

	if fromIndex < 0 || toIndex < 0 || fromIndex > len(column.Tasks)-1 || toIndex > len(column.Tasks)-1 {
		return ErrIndexOutOfBounds
	}

	task := column.Tasks[fromIndex]
	column.Tasks = slices.Delete(column.Tasks, fromIndex, fromIndex+1)
	column.Tasks = slices.Insert(column.Tasks, toIndex, task)

	return nil
}
