package cfgext_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nijeti/cfgext"
)

func TestWithStructTag(t *testing.T) {
	t.Parallel()

	const panicMsg = "invalid struct tag"

	tests := map[string]struct {
		tag    string
		panics bool
	}{
		"empty": {
			tag:    "",
			panics: true,
		},
		"invalid_semicolon": {
			tag:    "invalid;tag",
			panics: true,
		},
		"invalid_spaces": {
			tag:    "conf tag",
			panics: true,
		},
		"invalid_special_chars": {
			tag:    "conf@#$",
			panics: true,
		},
		"invalid_multiple_words": {
			tag:    "conf:tag,omitempty",
			panics: true,
		},
		"invalid_dots": {
			tag:    "conf.tag",
			panics: true,
		},
		"valid_simple": {
			tag:    "conf",
			panics: false,
		},
		"valid_complex": {
			tag:    "App-Conf_1",
			panics: false,
		},
	}

	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				params := cfgext.Params{}
				execFunc := func() { cfgext.WithStructTag(tt.tag)(&params) }

				if tt.panics {
					assert.PanicsWithValue(t, panicMsg, execFunc)
				} else {
					assert.NotPanics(t, execFunc)
					assert.Equal(t, params.StructTag, tt.tag)
				}
			},
		)
	}
}

func TestWithFilepath(t *testing.T) {
	t.Parallel()

	const panicMsg = "invalid filepath"

	tests := map[string]struct {
		filepath string
		panics   bool
	}{
		"empty": {
			filepath: "",
			panics:   true,
		},
		"invalid_extension": {
			filepath: "config.txt",
			panics:   true,
		},
		"invalid_no_extension": {
			filepath: "config",
			panics:   true,
		},
		"valid_file": {
			filepath: "config.yaml",
			panics:   false,
		},
		"valid_nested_file": {
			filepath: "config/config.yaml",
			panics:   false,
		},
	}

	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				params := cfgext.Params{}
				execFunc := func() { cfgext.WithFilepath(tt.filepath)(&params) }

				if tt.panics {
					assert.PanicsWithValue(t, panicMsg, execFunc)
				} else {
					assert.NotPanics(t, execFunc)
					assert.Equal(t, params.Filepath, tt.filepath)
				}
			},
		)
	}
}
