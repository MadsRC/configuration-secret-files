[![Go Report Card](https://goreportcard.com/badge/github.com/MadsRC/configuration-secret-files)](https://goreportcard.com/report/github.com/MadsRC/configuration-secret-files)

# configuration-secret-files

A provider for BoRuDar's [configuration](https://github.com/BoRuDar/configuration) package that loads configuration
values from files in a given directory.

On your configuration struct, you define the `secret_file` tag with the name of the file that contains the value you
want to load. The provider will read the file and provide the content of the file as the value for the field.

The provider will look for files in the `/run/secrets` directory by default. As such specifying the `secret_file` tag
with the value `my_secret` will make the provider look for a file named `/run/secrets/my_secret`.

The tag and the directory can be customized when creating the provider. Please refer to the
[documentation](https://pkg.go.dev/github.com/MadsRC/configuration-secret-files) for more information.
