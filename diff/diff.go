// Package diff строит дерево различий между двумя структурами — внутреннее представление для форматеров.
package diff

import (
	"maps"
	"reflect"
	"slices"
)

// Типы узлов дерева.
const (
	Root      = "root"
	Nested    = "nested"
	Added     = "added"
	Deleted   = "deleted"
	Changed   = "changed"
	Unchanged = "unchanged"
)

// Node — один ключ в дереве различий. Value1 — значение в первой структуре, Value2 — во второй;
// Children есть только у Nested и Root.
type Node struct {
	Key      string `json:"key"`
	Type     string `json:"type"`
	Value1   any    `json:"value1,omitempty"`
	Value2   any    `json:"value2,omitempty"`
	Children []Node `json:"children,omitempty"`
}

// Build возвращает корень дерева различий между data1 и data2.
func Build(data1, data2 map[string]any) Node {
	return Node{Type: Root, Children: buildChildren(data1, data2)}
}

// buildChildren описывает каждый ключ обеих структур, ключи — по алфавиту.
func buildChildren(data1, data2 map[string]any) []Node {
	keys := slices.Sorted(maps.Keys(data1))
	for key := range data2 {
		if _, ok := data1[key]; !ok {
			keys = append(keys, key)
		}
	}

	slices.Sort(keys)

	nodes := make([]Node, 0, len(keys))
	for _, key := range keys {
		nodes = append(nodes, buildNode(key, data1, data2))
	}

	return nodes
}

func buildNode(key string, data1, data2 map[string]any) Node {
	value1, in1 := data1[key]
	value2, in2 := data2[key]

	switch {
	case !in1:
		return Node{Key: key, Type: Added, Value2: value2}
	case !in2:
		return Node{Key: key, Type: Deleted, Value1: value1}
	}

	map1, isMap1 := value1.(map[string]any)
	map2, isMap2 := value2.(map[string]any)

	switch {
	case isMap1 && isMap2:
		return Node{Key: key, Type: Nested, Children: buildChildren(map1, map2)}
	case reflect.DeepEqual(value1, value2): // DeepEqual: в значениях бывают списки, их == роняет панику
		return Node{Key: key, Type: Unchanged, Value1: value1}
	default:
		return Node{Key: key, Type: Changed, Value1: value1, Value2: value2}
	}
}
