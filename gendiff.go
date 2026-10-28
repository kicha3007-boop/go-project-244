// Package code сравнивает два конфигурационных файла (JSON, YAML) и описывает различия в выбранном формате.
package code

import (
	"fmt"
	"os"
	"path/filepath"

	"code/diff"
	"code/formatters"
	"code/parsers"
)

// GenDiff возвращает различия между файлами path1 и path2 в формате format: stylish, plain или json.
// Пустой format означает stylish.
func GenDiff(path1, path2, format string) (string, error) {
	data1, err := readData(path1)
	if err != nil {
		return "", err
	}

	data2, err := readData(path2)
	if err != nil {
		return "", err
	}

	result, err := formatters.Format(diff.Build(data1, data2), format)
	if err != nil {
		return "", fmt.Errorf("format diff: %w", err)
	}

	return result, nil
}

// readData читает файл и разбирает его по расширению.
func readData(path string) (map[string]any, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}

	data, err := parsers.Parse(content, filepath.Ext(path))
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}

	return data, nil
}
