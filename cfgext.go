package cfgext

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

// Load initializes and returns a configuration of type T by loading data
// from specified sources such as files and environment variables. It accepts
// optional parameters to override default loading behavior.
func Load[T any](opts ...Option) (*T, error) {
	params := Params{
		StructTag: "conf",
		Filepath:  "config.yaml",
	}
	for _, opt := range opts {
		opt(&params)
	}

	k := koanf.New(".")

	if err := loadFile(k, params); err != nil {
		return nil, err
	}
	if err := loadEnv(k, params); err != nil {
		return nil, err
	}

	cfg := new(T)
	err := k.UnmarshalWithConf(
		"", cfg, koanf.UnmarshalConf{Tag: params.StructTag},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	return cfg, nil
}

func loadEnv(k *koanf.Koanf, _ Params) error {
	cb := func(s string) string {
		return strings.ReplaceAll(strings.ToLower(s), "__", ".")
	}
	err := k.Load(env.Provider("", ".", cb), nil)
	if err != nil {
		return fmt.Errorf("failed to load config from env: %w", err)
	}

	return nil
}

func loadFile(k *koanf.Koanf, params Params) error {
	err := k.Load(file.Provider(params.Filepath), yaml.Parser())
	switch {
	case err == nil:
		return nil
	case errors.Is(err, os.ErrNotExist):
		return nil
	default:
		return fmt.Errorf("failed to load config file: %w", err)
	}
}
