# Вычислитель отличий на Go

[![hexlet-check](https://github.com/kicha3007-boop/go-project-244/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/kicha3007-boop/go-project-244/actions)
[![CI](https://github.com/kicha3007-boop/go-project-244/actions/workflows/ci.yml/badge.svg)](https://github.com/kicha3007-boop/go-project-244/actions/workflows/ci.yml)

Консольная утилита и библиотека, которая сравнивает два конфигурационных файла (JSON, YAML) и показывает
различия в формате stylish, plain или json.

## Установка

Нужен Go 1.25 или новее.

```bash
git clone https://github.com/kicha3007-boop/go-project-244.git
cd go-project-244
make build              # бинарник в bin/gendiff
```

## Использование

```bash
./bin/gendiff --help
NAME:
   gendiff - Compares two configuration files and shows a difference.

USAGE:
   gendiff [global options] <filepath1> <filepath2>

GLOBAL OPTIONS:
   --format string, -f string  output format (default: "stylish")
   --help, -h                  show help
```

Формат входного файла определяется по расширению: `.json`, `.yml`, `.yaml`. Пути — относительные или абсолютные.

### stylish (по умолчанию)

```bash
./bin/gendiff testdata/fixture/file1.json testdata/fixture/file2.json
{
    common: {
      + follow: false
        setting1: Value 1
      - setting2: 200
      - setting3: true
      + setting3: null
...
```

### plain

```bash
./bin/gendiff --format plain testdata/fixture/file1.yml testdata/fixture/file2.yaml
Property 'common.follow' was added with value: false
Property 'common.setting2' was removed
Property 'common.setting3' was updated. From true to null
...
```

### json

```bash
./bin/gendiff -f json testdata/fixture/file1.json testdata/fixture/file2.json
```

Выводит дерево различий: у каждого узла `key`, `type` (`root`, `nested`, `added`, `deleted`, `changed`,
`unchanged`), значения `value1` / `value2` и `children` у вложенных.

## Как библиотека

```go
import "code"

result, err := code.GenDiff("file1.json", "file2.yml", "plain")
```

## Устройство

- `parsers` — разбор файла по расширению в `map[string]any`;
- `diff` — дерево различий (внутреннее представление), ключи по алфавиту;
- `formatters` — по файлу на формат, выбор формата в `formatters/formatters.go`.

## Разработка

```bash
make test            # тесты
make test-coverage   # покрытие, падает ниже COVERAGE_MIN из Makefile
make lint            # golangci-lint
```

Проверка Hexlet локально (Docker) — образ с тестами автопроверки, код кладётся в `/project/code`:

```bash
docker run --rm -v "$PWD:/src:ro" --entrypoint sh mirror.gcr.io/hexletprojects/go_gendiff_project_ru:latest \
  -c 'rm -rf /project/code && cp -r /src /project/code && cd /project && make setup && make test && make lint'
```
