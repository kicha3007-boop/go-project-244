package formatters

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"code/diff"
)

// indentSize — сколько пробелов добавляет каждый уровень вложенности.
const indentSize = 4

// Stylish выводит дерево фигурными скобками со знаками «+» и «-» у изменённых ключей.
func Stylish(tree diff.Node) string {
	return stylishBlock(tree.Children, 1)
}

// stylishBlock выводит детей узла глубины depth-1 в фигурных скобках.
func stylishBlock(nodes []diff.Node, depth int) string {
	lines := []string{"{"}
	for _, node := range nodes {
		lines = append(lines, stylishNode(node, depth)...)
	}

	lines = append(lines, strings.Repeat(" ", (depth-1)*indentSize)+"}")

	return strings.Join(lines, "\n")
}

func stylishNode(node diff.Node, depth int) []string {
	switch node.Type {
	case diff.Nested:
		return []string{stylishLine(depth, ' ', node.Key, stylishBlock(node.Children, depth+1))}
	case diff.Added:
		return []string{stylishLine(depth, '+', node.Key, stylishValue(node.Value2, depth))}
	case diff.Deleted:
		return []string{stylishLine(depth, '-', node.Key, stylishValue(node.Value1, depth))}
	case diff.Changed:
		return []string{
			stylishLine(depth, '-', node.Key, stylishValue(node.Value1, depth)),
			stylishLine(depth, '+', node.Key, stylishValue(node.Value2, depth)),
		}
	default:
		return []string{stylishLine(depth, ' ', node.Key, stylishValue(node.Value1, depth))}
	}
}

// stylishLine — строка ключа: отступ уровня минус два символа под знак и пробел.
func stylishLine(depth int, sign rune, key, value string) string {
	return fmt.Sprintf("%s%c %s: %s", strings.Repeat(" ", depth*indentSize-2), sign, key, value)
}

// stylishValue выводит значение как есть; вложенную структуру — блоком без знаков.
func stylishValue(value any, depth int) string {
	nested, ok := value.(map[string]any)
	if !ok {
		return scalar(value)
	}

	lines := []string{"{"}
	for _, key := range slices.Sorted(maps.Keys(nested)) {
		lines = append(lines, fmt.Sprintf("%s%s: %s",
			strings.Repeat(" ", (depth+1)*indentSize), key, stylishValue(nested[key], depth+1)))
	}

	lines = append(lines, strings.Repeat(" ", depth*indentSize)+"}")

	return strings.Join(lines, "\n")
}

// scalar выводит простое значение: null для отсутствующего, остальное — как есть.
func scalar(value any) string {
	if value == nil {
		return "null"
	}

	return fmt.Sprint(value)
}
