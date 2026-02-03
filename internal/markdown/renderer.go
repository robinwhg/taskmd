package markdown

func RenderColumn(column Column) string {
	return columnPrefix + column.Name
}
