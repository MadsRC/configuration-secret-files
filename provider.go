package configuration_secret_files

import (
	"fmt"
	"github.com/BoRuDar/configuration/v4"
	"io"
	"os"
	"path"
	"reflect"
)

const ProviderName = `SecretFilesProvider`

// NewProvider creates a new provider that reads secrets from files.
// The provider is configured with the specified options. A selection of options exists as functions that can be
// passed to this function. These functions are named `With*` where `*` is the name of the option.
// If no options are specified, the default options are used.
//
// Additionally, it is possible to set global options that will be applied to all new providers. These options
// overwrite the default options, but are overwritten by the options passed to this function.
//
// The default options are documented in the [Options] struct.
func NewProvider(options ...Option) configuration.Provider {
	opts := defaultOptions
	for _, opt := range GlobalOptions {
		opt.apply(&opts)
	}
	for _, opt := range options {
		opt.apply(&opts)
	}

	return &provider{
		options: &opts,
	}
}

type provider struct {
	options *Options
}

// Name returns the name of the provider.
func (p *provider) Name() string {
	return ProviderName
}

// Init initializes the provider.
// This is a no-op if [Options.DirectoryMustExist] is false. If true, it will check if the directory exists
// and return an error if it does not.
func (p *provider) Init(_ any) error {
	if p.options.DirectoryMustExist {
		if _, err := os.Stat(p.options.Directory); os.IsNotExist(err) {
			return fmt.Errorf("%s: directory '%s' does not exist: %w", ProviderName, p.options.Directory, err)
		}
	}
	return nil
}

// Provide reads the content of the file specified in the field tag and sets it to the field value.
// If the tag is not present, it is a no-op.
// If the tag is present, but the file does not exist or cannot be read, an error is returned.
// The file is read with a maximum size of whatever is defined with [Options.MaxSize]. If the file is
// larger, only the first [Options.MaxSize] bytes are read.
func (p *provider) Provide(field reflect.StructField, v reflect.Value) error {
	fileName, ok := field.Tag.Lookup(p.options.Tag)
	if !ok {
		return nil
	}

	filePath := path.Join(p.options.Directory, fileName)

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("%s: %w", ProviderName, err)
	}
	defer file.Close()

	content, err := io.ReadAll(io.LimitReader(file, p.options.MaxSize))
	if err != nil {
		return fmt.Errorf("%s: %w", ProviderName, err)
	}

	return configuration.SetField(field, v, string(content))
}
