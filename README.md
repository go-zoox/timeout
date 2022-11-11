# Timeout - Give timeout power for function

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/timeout)](https://pkg.go.dev/github.com/go-zoox/timeout)
[![Build Status](https://github.com/go-zoox/timeout/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/timeout/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/timeout)](https://goreportcard.com/report/github.com/go-zoox/timeout)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/timeout/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/timeout?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/timeout.svg)](https://github.com/go-zoox/timeout/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/timeout.svg?label=Release)](https://github.com/go-zoox/timeout/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/timeout
```

## Getting Started

```go
func TestTimeout(t *testing.T) {
	fn := func() error {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("fn done")
		return nil
	}

	err := Timeout(fn, 50*time.Millisecond)
	if err == nil {
		t.Errorf("expected timeout error, got nil")
	}

	err = Timeout(fn, 300*time.Millisecond)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}
```

## License
GoZoox is released under the [MIT License](./LICENSE).
