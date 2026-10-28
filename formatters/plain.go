package formatters

import (
	"fmt"
	"strings"

	"code/diff"
)

// Plain выводит по строке на каждое изменённое свойство с полным путём до него.
func Plain(tree diff.Node) string {
	return strings.Join(plainLines(tree.Children, ""), "\n")
}

func plainLines(nodes []diff.Node, parent string) []string {
	var lines []string

	for _, node := range nodes {
		path := node.Key
		if parent != "" {
			path = parent + "." + node.Key
		}

		switch node.Type {
		case diff.Nested:
			lines = append(lines, plainLines(node.Children, path)...)
		case diff.Added:
			lines = append(lines, fmt.Sprintf("Property '%s' was added with value: %s", path, plainValue(node.Value2)))
		case diff.Deleted:
			lines = append(lines, fmt.Sprintf("Property '%s' was removed", path))
		case diff.Changed:
			lines = append(lines, fmt.Sprintf("Property '%s' was updated. From %s to %s",
				path, plainValue(node.Value1), plainValue(node.Value2)))
		}
	}

	return lines
}

// plainValue: строки в одинарных кавычках, составные значения — [complex value].
func plainValue(value any) string {
	switch v := value.(type) {
	case map[string]any:
		return "[complex value]"
	case string:
		return "'" + v + "'"
	default:
		return scalar(v)
	}
}
