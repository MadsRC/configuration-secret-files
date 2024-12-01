package configuration_secret_files

// Options is the configuration for the provider.
// They are not used directly, but are passed to the provider when it is created with [NewProvider], via the [Option]
// interface.
type Options struct {
	// Directory is the directory where the secret files are located. Default is "/run/secrets".
	Directory string
	// Tag is the tag that is used to specify the file name in the struct field. Default is "secret_file".
	Tag string
	// MaxSize is the maximum size of the file that is read. Default is 4MiB.
	MaxSize int64
	// DirectoryMustExist specifies if the directory must exist. Default is true.
	DirectoryMustExist bool
}

var defaultOptions = Options{
	Directory:          "/run/secrets",
	Tag:                "secret_file",
	MaxSize:            4 * 1024 * 1024,
	DirectoryMustExist: true,
}

// GlobalOptions is a list of [Options] that will be applied to all new providers.
var GlobalOptions []Option

// Option is an option for configuring a [Provider].
type Option interface {
	apply(*Options)
}

// funcOption is an Option that calls a function.
// It is used to wrap a function, so it satisfies the Option interface.
type funcOption struct {
	f func(*Options)
}

func (fdo *funcOption) apply(opts *Options) {
	fdo.f(opts)
}

func newFuncOption(f func(*Options)) *funcOption {
	return &funcOption{
		f: f,
	}
}

// WithDirectory sets the [Options].Directory to the specified value.
func WithDirectory(directory string) Option {
	return newFuncOption(func(opts *Options) {
		opts.Directory = directory
	})
}

// WithTag sets the [Options].Tag to the specified value.
func WithTag(tag string) Option {
	return newFuncOption(func(opts *Options) {
		opts.Tag = tag
	})
}

// WithMaxSize sets the [Options].MaxSize to the specified value.
func WithMaxSize(maxSize int64) Option {
	return newFuncOption(func(opts *Options) {
		opts.MaxSize = maxSize
	})
}

// WithDirectoryMustExist sets the [Options].DirectoryMustExist to the specified value.
func WithDirectoryMustExist(mustExist bool) Option {
	return newFuncOption(func(opts *Options) {
		opts.DirectoryMustExist = mustExist
	})
}
