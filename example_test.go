package configuration_secret_files_test

import (
	"fmt"
	"github.com/BoRuDar/configuration/v4"
	configuration_secret_files "github.com/MadsRC/configuration-secret-files"
	"os"
	"path/filepath"
)

const Payload = `Hello, World!`

var dir string

func SetupExample() (func(), error) {
	var err error
	dir, err = os.MkdirTemp("", "example")
	if err != nil {
		return nil, err
	}
	fpath := filepath.Join(dir, "secret")
	err = os.WriteFile(fpath, []byte(Payload), 0644)
	if err != nil {
		return nil, err
	}

	return func() {
		os.RemoveAll(dir)
	}, nil
}

type Config struct {
	SomethingNotSecret string
	TheSecret          string `secret_file:"secret"`
}

func Example() {
	// IGNORE THIS PART OF THE EXAMPLE, WE NEED TO DO SOME SETUP
	cleanup, err := SetupExample()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cleanup()
	// END OF THE SETUP - STOP IGNORING

	cfg := &Config{}

	configurator := configuration.New(cfg, configuration_secret_files.NewProvider(
		// This is for the sake of the example. If we don't set the directory, the default one will be used.
		configuration_secret_files.WithDirectory(filepath.Join(dir)),
	))

	err = configurator.InitValues()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", cfg)

	// Output:
	// &{SomethingNotSecret: TheSecret:Hello, World!}
}
