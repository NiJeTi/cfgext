package cfgext

import (
	"strings"
)

// Params represents options for configuration loading.
// StructTag specifies the struct tag used during unmarshaling (default: conf).
// Filepath specifies the path to the configuration file (default: config.yaml).
type Params struct {
	StructTag string
	Filepath  string
}

// Option defines a function that modifies Params configuration.
type Option func(params *Params)

// WithStructTag sets the struct tag to be used during unmarshaling.
// Panics if the provided tag is invalid.
func WithStructTag(tag string) Option {
	if !isValidStructTag(tag) {
		panic("invalid struct tag")
	}

	return func(conf *Params) {
		conf.StructTag = tag
	}
}

// WithFilepath sets the file path for the configuration file in Params.
// Panics if the provided file path does not have a ".yaml" suffix.
func WithFilepath(filepath string) Option {
	if !strings.HasSuffix(filepath, ".yaml") {
		panic("invalid filepath")
	}

	return func(conf *Params) {
		conf.Filepath = filepath
	}
}

func isValidStructTag(tag string) bool {
	if len(tag) == 0 {
		return false
	}

	for _, c := range tag {
		if c >= 'a' && c <= 'z' {
			continue
		}
		if c >= 'A' && c <= 'Z' {
			continue
		}
		if c >= '0' && c <= '9' {
			continue
		}
		if c == '-' || c == '_' {
			continue
		}

		return false
	}

	return true
}
