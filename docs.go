// This package provides a provider for [github.com/BoRuDar/configuration] that reads configuration values
// from files.
//
// Each field in the configuration struct can be tagged with a file name. The provider reads the content of the file
// and sets the field value to the content of the file.
//
// The main use-case is to provide a way to read sensitive/secret values from files, as an alternative to providing them
// as environment variables or command-line arguments. This is especially useful in containerized environments, where
// the secret values can be mounted as files. See the Docker documentation at
// https://docs.docker.com/engine/swarm/secrets/ or the Kubernetes documentation at
// https://kubernetes.io/docs/concepts/configuration/secret/#using-secrets-as-files-from-a-pod for more information
// on this concept.
//
// A new provider is created with the [NewProvider] function. The provider is configured with options that are passed
// as functions to the [NewProvider] function.
//
// # Options
//
// Options used to configured the provider can be set globally or per provider. Global options are applied to all new
// providers. They can be set with the [GlobalOptions] variable. Options set per provider overwrite the global options.
//
// A selection of functions that configure the provider are available. These functions are named `With*`.
//
// Additionally, the documentation for the [Options] struct provides information about the available options and
// their default values.
//
// # Security of the secret files
//
// The nature of this provider inevitably means that your secrets will be written in plain text to your filesystem.
// This is intentional, albeit not ideal. It is meant as an alternative to providing secrets as environment variables
// or command-line arguments. The security of the secret files is the responsibility of the user of this provider.
//
// Please make sure to assess the environment in which you plan to handle secrets, and only use this provider if you
// are confident that writing the secrets to the filesystem is secure more secure than other methods.
package configuration_secret_files
