package cfgext_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nijeti/cfgext"
)

func TestLoad(t *testing.T) {
	t.Parallel()

	type nestedCfg struct {
		Name string `conf:"name"`
	}
	type config struct {
		Name    string    `conf:"name"`
		Version string    `conf:"version"`
		Nested  nestedCfg `conf:"nested"`
	}

	tests := map[string]struct {
		options []cfgext.Option
		cfg     *config
		isErr   bool
		setup   func()
	}{
		"file_not_found": {
			options: []cfgext.Option{
				cfgext.WithFilepath("testdata/missing.yaml"),
			},
			cfg:   &config{},
			isErr: false,
		},
		"invalid_file": {
			options: []cfgext.Option{
				cfgext.WithFilepath("testdata/invalid.yaml"),
			},
			cfg:   nil,
			isErr: true,
			setup: func() {
				os.MkdirAll("testdata", 0755)
				os.WriteFile(
					"testdata/invalid.yaml",
					[]byte("invalid_yaml_content"),
					0644,
				)
			},
		},
		"valid_file": {
			options: []cfgext.Option{
				cfgext.WithFilepath("testdata/config.yaml"),
			},
			cfg:   &config{Name: "test-app", Version: "1.0.0"},
			isErr: false,
			setup: func() {
				os.MkdirAll("testdata", 0755)
				os.WriteFile(
					"testdata/config.yaml",
					[]byte("name: test-app\nversion: 1.0.0"),
					0644,
				)
			},
		},
		"valid_env": {
			options: []cfgext.Option{},
			setup: func() {
				os.Setenv("NAME", "test-app")
				os.Setenv("VERSION", "1.0.0")
			},
			cfg: &config{Name: "test-app", Version: "1.0.0"},
		},
		"nested_env": {
			options: []cfgext.Option{},
			setup: func() {
				os.Setenv("NESTED__NAME", "nested-name")
			},
			cfg: &config{Nested: nestedCfg{Name: "nested-name"}},
		},
		"empty_config": {
			options: []cfgext.Option{},
			cfg:     &config{},
			isErr:   false,
		},
	}

	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				defer os.RemoveAll("testdata")

				defer os.Unsetenv("NAME")
				defer os.Unsetenv("VERSION")
				defer os.Unsetenv("NESTED__NAME")

				if tt.setup != nil {
					tt.setup()
				}

				cfg, err := cfgext.Load[config](tt.options...)

				assert.Equal(t, tt.cfg, cfg)
				if tt.isErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
			},
		)
	}
}
