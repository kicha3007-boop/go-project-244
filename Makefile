BINARY := bin/gendiff
# пакеты под порогом покрытия (их код проверяют и тесты других пакетов); main из-под порога исключён
COVER_PKGS := code,code/diff,code/formatters,code/parsers
COVERAGE_MIN := 85

.PHONY: build run test test-coverage lint lint-fix check

build:
	go build -o $(BINARY) ./cmd/gendiff

run: build
	./$(BINARY) $(ARGS)

test:
	go test -v ./...

test-coverage:
	go test -coverpkg=$(COVER_PKGS) -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out | awk -v min=$(COVERAGE_MIN) '/^total:/ { sub("%", "", $$3); \
		printf "coverage %s%% (min %s%%)\n", $$3, min; if ($$3 + 0 < min) { print "coverage below threshold"; exit 1 } }'

lint:
	golangci-lint run ./...

lint-fix:
	golangci-lint run --fix ./...

check: lint test-coverage
