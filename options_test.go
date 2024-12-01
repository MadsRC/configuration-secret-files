package configuration_secret_files

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWithDirectory(t *testing.T) {
	t.Parallel()
	cfg := &Options{}
	WithDirectory("/a/path").apply(cfg)
	require.Equal(t, "/a/path", cfg.Directory)
}

func TestWithTag(t *testing.T) {
	t.Parallel()
	cfg := &Options{}
	WithTag("tag").apply(cfg)
	require.Equal(t, "tag", cfg.Tag)
}

func TestWithMaxSize(t *testing.T) {
	t.Parallel()
	cfg := &Options{}
	WithMaxSize(42).apply(cfg)
	require.Equal(t, int64(42), cfg.MaxSize)
}

func TestWithDirectoryMustExist(t *testing.T) {
	t.Parallel()
	cfg := &Options{}
	WithDirectoryMustExist(true).apply(cfg)
	require.Equal(t, true, cfg.DirectoryMustExist)
}

func TestFuncOption(t *testing.T) {
	t.Parallel()
	cfg := &Options{}
	newFuncOption(func(opts *Options) {
		opts.Directory = "/a/path"
	}).apply(cfg)
	require.Equal(t, "/a/path", cfg.Directory)
}
