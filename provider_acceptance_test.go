package configuration_secret_files_test

import (
	"github.com/BoRuDar/configuration/v4"
	configuration_secret_files "github.com/MadsRC/configuration-secret-files"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

// TestProvider_simple tests the provider with a simple configuration.
// The configuration has a field that is tagged with a file name.
// A file with the specified name is read from the directory and the content is set to the field.
//
// The configuration has a field that is not tagged with a file name. The provider
// should ignore the field and not set any value or return an error.
func TestProvider_simple(t *testing.T) {
	t.Parallel()
	expected, err := os.ReadFile("testdata/secretfile")
	require.NoError(t, err)
	p := configuration_secret_files.NewProvider(
		configuration_secret_files.WithDirectory("testdata"),
	)

	type testCfg struct {
		Field  string `secret_file:"secretfile"`
		NotSet string
	}
	cfg := &testCfg{}

	configurator := configuration.New(cfg, p)
	err = configurator.InitValues()
	require.NoError(t, err)
	require.Equal(t, string(expected), cfg.Field)
	require.Equal(t, "", cfg.NotSet)
}
