// Package markdown handles the parsing of markdown content
package markdown

import (
	"bufio"
	"io"
	"strings"
)

const (
	columnPrefix      = "## "
	taskPrefix        = "- [ ] "
	checkedTaskPrefix = "- [x] "
)

type Task struct {
	Name    string
	Checked bool
}

type Column struct {
	Name  string
	Tasks []Task
}

type Board struct {
	Columns []Column
}

func Parse(input io.Reader) (*Board, error) {
	board := Board{}

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		text := scanner.Text()

		column, found := parseColumn(text)
		if found {
			board.Columns = append(board.Columns, column)
			continue
		}

		task, found := parseTask(text)
		if found {
			currentBoard := &board.Columns[len(board.Columns)-1] // FIXME: a column has to exist
			currentBoard.Tasks = append(currentBoard.Tasks, task)
		}
	}

	return &board, nil
}

func parseColumn(text string) (column Column, found bool) {
	after, found := strings.CutPrefix(text, columnPrefix)

	if found {
		return Column{Name: after}, true
	}

	return Column{}, false
}

func parseTask(text string) (task Task, found bool) {
	after, found := strings.CutPrefix(text, taskPrefix)
	if found {
		return Task{Name: after, Checked: false}, true
	}

	after, found = strings.CutPrefix(text, checkedTaskPrefix)
	if found {
		return Task{Name: after, Checked: true}, true
	}

	return Task{}, false
}
