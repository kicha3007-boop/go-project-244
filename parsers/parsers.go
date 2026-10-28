// Package parsers разбирает содержимое конфигурационных файлов в map[string]any.
package parsers

import (
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// Parse разбирает content по расширению файла ext: .json, .yml или .yaml.
func Parse(content []byte, ext string) (map[string]any, error) {
	data := map[string]any{}

	var err error

	switch strings.ToLower(ext) {
	case ".json":
		err = json.Unmarshal(content, &data)
	case ".yml", ".yaml":
		err = yaml.Unmarshal(content, &data)
	default:
		return nil, fmt.Errorf("unsupported file format %q", ext)
	}

	if err != nil {
		return nil, fmt.Errorf("decode %s: %w", ext, err)
	}

	return data, nil
}
