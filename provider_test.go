package configuration_secret_files

import (
	"github.com/stretchr/testify/require"
	"os"
	"reflect"
	"testing"
)

const dirDoesNotExist = "/this/dir/does/not/exist"

func TestNewProvider(t *testing.T) {
	t.Parallel()
	p := NewProvider()
	require.NotNil(t, p)
}

func TestNewProvider_defaults_are_set(t *testing.T) {
	t.Parallel()
	p := NewProvider()
	require.NotNil(t, p)

	pr, ok := p.(*provider)
	require.True(t, ok)
	require.NotNil(t, pr.options)
	require.Equal(t, defaultOptions, *pr.options)
}

func TestNewProvider_global_options_are_applied(t *testing.T) {
	GlobalOptions = append(GlobalOptions, WithDirectory("/a/path"))
	defer func() {
		GlobalOptions = GlobalOptions[:len(GlobalOptions)-1]
	}()

	p := NewProvider()
	require.NotNil(t, p)

	pr, ok := p.(*provider)
	require.True(t, ok)
	require.NotNil(t, pr.options)
	require.Equal(t, "/a/path", pr.options.Directory)
}

func TestNewProvider_options_are_applied(t *testing.T) {
	GlobalOptions = append(GlobalOptions, WithDirectory("/another/path"))
	defer func() {
		GlobalOptions = GlobalOptions[:len(GlobalOptions)-1]
	}()
	p := NewProvider(WithDirectory("/a/path"))
	require.NotNil(t, p)

	pr, ok := p.(*provider)
	require.True(t, ok)
	require.NotNil(t, pr.options)
	require.NotEqual(t, "/another/path", pr.options.Directory, "global options have been overwritten")
	require.Equal(t, "/a/path", pr.options.Directory)
}

func TestProvider_Name(t *testing.T) {
	t.Parallel()
	p := provider{}
	require.Equal(t, ProviderName, p.Name())
}

func TestProvider_Init_no_directory_must_exist(t *testing.T) {
	t.Parallel()
	p := provider{
		options: &Options{
			Directory:          dirDoesNotExist,
			DirectoryMustExist: true,
		},
	}
	require.Error(t, p.Init(nil))
}

func TestProvider_Init_directory(m *testing.T) {
	m.Parallel()
	p := provider{
		options: &Options{
			Directory:          dirDoesNotExist,
			DirectoryMustExist: false,
		},
	}
	require.NoError(m, p.Init(nil))
}

func TestProvider_Init_directory_exists(t *testing.T) {
	cwd, err := os.Getwd()
	require.NoError(t, err)
	t.Parallel()
	p := provider{
		options: &Options{
			Directory:          cwd,
			DirectoryMustExist: true,
		},
	}
	require.NoError(t, p.Init(nil))
}

func TestProvider_Provide_no_tag(t *testing.T) {
	t.Parallel()
	p := provider{
		options: &Options{
			Tag: "tag",
		},
	}
	type something struct {
		NoTag string
	}

	cfg := something{}
	require.NoError(t, p.Provide(reflect.TypeOf(cfg).Field(0), reflect.ValueOf(&cfg).Elem().Field(0)))
	require.Empty(t, cfg.NoTag)
}

func TestProvider_Provide_file_does_not_exist(t *testing.T) {
	t.Parallel()
	p := provider{
		options: &Options{
			Tag: "tag",
		},
	}
	type something struct {
		Tag string `tag:"file_does_not_exist"`
	}

	cfg := something{}
	require.Error(t, p.Provide(reflect.TypeOf(cfg).Field(0), reflect.ValueOf(&cfg).Elem().Field(0)))
}

func TestProvider_Provide_file_exists(t *testing.T) {
	t.Parallel()
	p := provider{
		options: &Options{
			Directory: "testdata",
			Tag:       "tag",
			MaxSize:   4 * 1024 * 1024,
		},
	}
	type something struct {
		Tag string `tag:"secretfile"`
	}

	cfg := something{}
	require.NoError(t, p.Provide(reflect.TypeOf(cfg).Field(0), reflect.ValueOf(&cfg).Elem().Field(0)))
	require.NotEmpty(t, cfg.Tag)
}

func TestProvider_Provide_file_too_large(t *testing.T) {
	t.Parallel()
	p := provider{
		options: &Options{
			Directory: "testdata",
			Tag:       "tag",
			MaxSize:   4,
		},
	}
	type something struct {
		Tag string `tag:"secretfile"`
	}

	cfg := something{}
	require.NoError(t, p.Provide(reflect.TypeOf(cfg).Field(0), reflect.ValueOf(&cfg).Elem().Field(0)))
	require.Equal(t, "This", cfg.Tag)
}
