package markdown

func RenderColumn(column Column) string {
	return columnPrefix + column.Name
}

func RenderTask(task Task) string {
	if task.Checked {
		return checkedTaskPrefix + task.Name
	} else {
		return taskPrefix + task.Name
	}
}
