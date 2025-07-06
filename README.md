# cfgext

[![Coverage Status](https://coveralls.io/repos/github/NiJeTi/cfgext/badge.svg)](https://coveralls.io/github/NiJeTi/cfgext)

A small extension/wrapper for [koanf](https://github.com/knadh/koanf) package
to simplify both deploy and local development.

Package supports two ways of loading config in the following priority:
1. YAML config file
2. Environment variables

## Installation

```shell
go get github.com/nijeti/cfgext
```

## Usage

Define your config structures (package supports nesting) and tag them with `conf` tag.

```go
type nestedConf struct {
	Var1 string `conf:"var1"`
}

type rootConf struct {
	Nested nestedConf `conf:"nested"`
	Var1   string     `conf:"var1"`
	Var2   int        `conf:"var2"`
}
```

Specify parameters either in environment variables or config file.

```shell
VAR1="var1_value"
VAR2="123"
NESTED__VAR1="nested_var1_value"
```

```yaml
var1: 'var1_value'
var2: 123
nested:
  var1: 'nested_var1_value'
```

Load config with `Load` function.

#### With default options
```go
cfg, err := cfgext.Load[rootConf]()
if err != nil {
	panic("failed to load config")
}
```

#### With custom options
```go
cfg, err := cfgext.Load[rootConf](
	cfgext.WithStructTag("conf"),
	cfgext.WithFilepath("config/app.yaml"),
)
if err != nil {
	panic("failed to load config")
}
```

## License

This project is licensed under the MIT License.
