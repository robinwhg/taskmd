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
	Line    int
}

type Column struct {
	Name  string
	Tasks []Task
	Line  int
}

type Board struct {
	Title   string
	Columns []Column
}

func Parse(input io.Reader) (*Board, error) {
	board := Board{}

	scanner := bufio.NewScanner(input)

	lineNum := -1
	for scanner.Scan() {
		lineNum++
		text := scanner.Text()

		if task, found := parseTask(text, lineNum); found {
			if len(board.Columns) > 0 {
				currentBoard := &board.Columns[len(board.Columns)-1]
				currentBoard.Tasks = append(currentBoard.Tasks, task)
				continue
			}

			board.Columns = append(board.Columns, Column{Name: "Uncategorized", Tasks: []Task{task}})
			continue
		}

		if column, found := parseColumn(text, lineNum); found {
			board.Columns = append(board.Columns, column)
			continue
		}

		if board.Title == "" && len(board.Columns) == 0 {
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
		return strings.TrimSpace(after), true
	}

	return "", false
}

func parseColumn(text string, lineNum int) (column Column, found bool) {
	if after, found := strings.CutPrefix(text, columnPrefix); found {
		return Column{Name: strings.TrimSpace(after), Line: lineNum}, true
	}

	return Column{}, false
}

func parseTask(text string, lineNum int) (task Task, found bool) {
	if after, found := strings.CutPrefix(text, taskPrefix); found {
		return Task{Name: strings.TrimSpace(after), Checked: false, Line: lineNum}, true
	}

	if after, found := strings.CutPrefix(text, checkedTaskPrefix); found {
		return Task{Name: strings.TrimSpace(after), Checked: true, Line: lineNum}, true
	}

	return Task{}, false
}
