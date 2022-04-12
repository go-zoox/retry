# Retry - Catch Panic In Retries

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/retry)](https://pkg.go.dev/github.com/go-zoox/retry)
[![Build Status](https://github.com/go-zoox/retry/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/retry/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/retry)](https://goreportcard.com/report/github.com/go-zoox/retry)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/retry/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/retry?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/retry.svg)](https://github.com/go-zoox/retry/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/retry.svg?label=Release)](https://github.com/go-zoox/retry/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/retry
```

## Getting Started

```go
func TestMemoryretry(t *testing.T) {
	m := retry.NewMemory()
	if m.Size() != 0 {
		t.Errorf("Expected size 0, got %d", m.Size())
	}

	m.Set("key", "value")
	if m.Get("key") != "value" {
		t.Error("Expected value to be 'value'")
	}

	if m.Size() != 1 {
		t.Errorf("Expected size 1, got %d", m.Size())
	}
}
```

## Engines
* [x] Memory
* [x] Redis
* [ ] MongoDB
* [x] SQLite
* [ ] PostgreSQL
* [ ] MySQL
* [ ] DynamoDB

## Inspired by
* [srfrog/dict](https://github.com/srfrog/dict) - Python-like dictionaries for Go

## License
GoZoox is released under the [MIT License](./LICENSE).