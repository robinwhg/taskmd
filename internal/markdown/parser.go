// Package markdown handles the parsing of markdown content
package markdown

import (
	"bufio"
	"io"
	"strings"
)

const (
	titlePrefix       = "# "
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
	Title   string
	Columns []Column
}

func Parse(input io.Reader) (*Board, error) {
	board := Board{}

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		text := scanner.Text()

		if task, found := parseTask(text); found {
			currentBoard := &board.Columns[len(board.Columns)-1] // FIXME: a column has to exist
			currentBoard.Tasks = append(currentBoard.Tasks, task)
			continue
		}

		if column, found := parseColumn(text); found {
			board.Columns = append(board.Columns, column)
			continue
		}

		if board.Title == "" {
			if title, found := parseTitle(text); found {
				board.Title = title
				continue
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &board, nil
}

func parseTitle(text string) (title string, found bool) {
	if after, found := strings.CutPrefix(text, titlePrefix); found {
		return after, true
	}

	return "", false
}

func parseColumn(text string) (column Column, found bool) {
	if after, found := strings.CutPrefix(text, columnPrefix); found {
		return Column{Name: after}, true
	}

	return Column{}, false
}

func parseTask(text string) (task Task, found bool) {
	if after, found := strings.CutPrefix(text, taskPrefix); found {
		return Task{Name: after, Checked: false}, true
	}

	if after, found := strings.CutPrefix(text, checkedTaskPrefix); found {
		return Task{Name: after, Checked: true}, true
	}

	return Task{}, false
}
