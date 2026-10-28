package diff

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuild(t *testing.T) {
	data1 := map[string]any{
		"same":    1,
		"changed": "a",
		"deleted": true,
		"list":    []any{1, 2},
		"nested":  map[string]any{"x": 1},
		"toMap":   "str",
	}
	data2 := map[string]any{
		"same":    1,
		"changed": "b",
		"added":   nil,
		"list":    []any{1, 2},
		"nested":  map[string]any{"x": 2},
		"toMap":   map[string]any{"k": "v"},
	}

	tree := Build(data1, data2)
	require.Equal(t, Root, tree.Type)

	expected := []Node{
		{Key: "added", Type: Added},
		{Key: "changed", Type: Changed, Value1: "a", Value2: "b"},
		{Key: "deleted", Type: Deleted, Value1: true},
		{Key: "list", Type: Unchanged, Value1: []any{1, 2}},
		{Key: "nested", Type: Nested, Children: []Node{{Key: "x", Type: Changed, Value1: 1, Value2: 2}}},
		{Key: "same", Type: Unchanged, Value1: 1},
		{Key: "toMap", Type: Changed, Value1: "str", Value2: map[string]any{"k": "v"}},
	}
	require.Equal(t, expected, tree.Children)
}
