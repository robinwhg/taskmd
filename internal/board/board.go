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

func DeleteTask(column *markdown.Column, index int) error {
	if column == nil {
		return ErrNilColumn
	}
	if err := checkIndexOutOfBounds(column, index); err != nil {
		return err
	}
	column.Tasks = slices.Delete(column.Tasks, index, index+1)
	return nil
}

func MoveTaskInColumn(column *markdown.Column, fromIndex, toIndex int) error {
	if column == nil {
		return ErrNilColumn
	}
	if fromIndex == toIndex {
		return nil
	}
	if err := checkIndexOutOfBounds(column, toIndex, fromIndex); err != nil {
		return err
	}

	task := column.Tasks[fromIndex]
	column.Tasks = slices.Delete(column.Tasks, fromIndex, fromIndex+1)
	column.Tasks = slices.Insert(column.Tasks, toIndex, task)

	return nil
}

func MoveTaskToColumn(fromColumn *markdown.Column, fromIndex int, toColumn *markdown.Column, toIndex int) error {
	if fromColumn == nil || toColumn == nil {
		return ErrNilColumn
	}
	if err := checkIndexOutOfBounds(fromColumn, fromIndex); err != nil {
		return err
	}
	if toIndex < 0 || toIndex > len(toColumn.Tasks) {
		return ErrIndexOutOfBounds
	}

	task := fromColumn.Tasks[fromIndex]
	fromColumn.Tasks = slices.Delete(fromColumn.Tasks, fromIndex, fromIndex+1)
	toColumn.Tasks = slices.Insert(toColumn.Tasks, toIndex, task)

	return nil
}

func checkIndexOutOfBounds(column *markdown.Column, index ...int) error {
	for _, i := range index {
		if i < 0 || i > len(column.Tasks)-1 {
			return ErrIndexOutOfBounds
		}
	}
	return nil
}
