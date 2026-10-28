package code

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func fixturePath(name string) string {
	return filepath.Join("testdata", "fixture", name)
}

func readFixture(t *testing.T, name string) string {
	t.Helper()

	content, err := os.ReadFile(fixturePath(name))
	require.NoError(t, err)

	return strings.TrimSuffix(string(content), "\n")
}

func TestGenDiff(t *testing.T) {
	inputs := []struct{ file1, file2 string }{
		{"file1.json", "file2.json"},
		{"file1.yml", "file2.yaml"},
		{"file1.json", "file2.yaml"},
	}
	formats := []struct{ format, expected string }{
		{"stylish", "result_stylish.txt"},
		{"", "result_stylish.txt"},
		{"plain", "result_plain.txt"},
	}

	for _, in := range inputs {
		for _, f := range formats {
			t.Run(in.file1+" "+in.file2+" "+f.format, func(t *testing.T) {
				got, err := GenDiff(fixturePath(in.file1), fixturePath(in.file2), f.format)
				require.NoError(t, err)
				require.Equal(t, readFixture(t, f.expected), got)
			})
		}
	}
}

func TestGenDiffJSON(t *testing.T) {
	got, err := GenDiff(fixturePath("file1.json"), fixturePath("file2.yaml"), "json")
	require.NoError(t, err)

	var tree struct {
		Type     string `json:"type"`
		Children []struct {
			Key  string `json:"key"`
			Type string `json:"type"`
		} `json:"children"`
	}
	require.NoError(t, json.Unmarshal([]byte(got), &tree))
	require.Equal(t, "root", tree.Type)

	types := map[string]string{}
	for _, child := range tree.Children {
		types[child.Key] = child.Type
	}

	require.Equal(t, map[string]string{
		"common": "nested", "group1": "nested", "group2": "deleted", "group3": "added",
	}, types)
}

func TestGenDiffErrors(t *testing.T) {
	_, err := GenDiff(fixturePath("missing.json"), fixturePath("file2.json"), "stylish")
	require.ErrorContains(t, err, "missing.json")

	_, err = GenDiff(fixturePath("file1.json"), fixturePath("result_plain.txt"), "stylish")
	require.ErrorContains(t, err, "unsupported file format")

	_, err = GenDiff(fixturePath("file1.json"), fixturePath("file2.json"), "xml")
	require.ErrorContains(t, err, "unknown format")
}
