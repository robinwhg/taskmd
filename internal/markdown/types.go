package markdown

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
