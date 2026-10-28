// Package formatters переводит дерево различий в текст: stylish, plain или json.
package formatters

import (
	"fmt"

	"code/diff"
)

// Format выводит дерево в формате format; пустой format означает stylish.
func Format(tree diff.Node, format string) (string, error) {
	switch format {
	case "", "stylish":
		return Stylish(tree), nil
	case "plain":
		return Plain(tree), nil
	case "json":
		return JSON(tree)
	default:
		return "", fmt.Errorf("unknown format %q", format)
	}
}
