# configuration-secret-files

This repository contains a GoLang module providing a provider for BoRuDar's
[configuration](https://github.com/BoRuDar/configuration) package.

The provider loads individual configuration values from files in a given directory. This is useful for loading
secret or sensitive values into your configuration when running Docker Swarm or Kubernetes.
