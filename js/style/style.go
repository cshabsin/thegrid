package style

type Style struct {
	Name  string
	Value string
}

func Make(name string, value string) Style {
	return Style{Name: name, Value: value}
}

func Border(value string) Style {
	return Make("border", value)
}

func BackgroundColor(value string) Style {
	return Make("background-color", value)
}

func Color(value string) Style {
	return Make("color", value)
}

func Position(value string) Style {
	return Make("position", value)
}

func Top(value string) Style {
	return Make("top", value)
}

func Left(value string) Style {
	return Make("left", value)
}

func GridColumn(value string) Style {
	return Make("grid-column", value)
}

func GridRow(value string) Style {
	return Make("grid-row", value)
}

func Display(value string) Style {
	return Make("display", value)
}

func GridTemplateColumns(value string) Style {
	return Make("grid-template-columns", value)
}

func FontSize(value string) Style {
	return Make("font-size", value)
}

func Transform(value string) Style {
	return Make("transform", value)
}

func Height(value string) Style {
	return Make("height", value)
}

func Width(value string) Style {
	return Make("width", value)
}
