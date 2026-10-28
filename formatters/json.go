package formatters

import (
	"encoding/json"
	"fmt"

	"code/diff"
)

// JSON выводит дерево различий как есть — для других программ.
func JSON(tree diff.Node) (string, error) {
	out, err := json.MarshalIndent(tree, "", "  ")
	if err != nil {
		return "", fmt.Errorf("marshal diff: %w", err)
	}

	return string(out), nil
}
