# Binbrew

[![Go Version](https://img.shields.io/badge/go%20version-%3E=1.11-orange.svg?style=flat-square)](https://github.com/sagikazarmark/binbrew)
[![CircleCI](https://circleci.com/gh/sagikazarmark/binbrew.svg?style=svg)](https://circleci.com/gh/sagikazarmark/binbrew)
[![Go Report Card](https://goreportcard.com/badge/github.com/sagikazarmark/binbrew?style=flat-square)](https://goreportcard.com/report/github.com/sagikazarmark/binbrew)
[![GolangCI](https://golangci.com/badges/github.com/sagikazarmark/binbrew.svg)](https://golangci.com/r/github.com/sagikazarmark/binbrew)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/sagikazarmark/binbrew)

Binbrew installs pre-built binary dependencies (from Github).
It's primary use case is to help setting up development environments and
allow to use the same tools in CI environments as well.


## Quick start

In order to install binaries run the following command:

```bash
binbrew install dep@0.5.0
```

You can install multiple binaries with a single command:

```bash
binbrew install dep@0.5.0 protoc@3.6.1
```


## Usage

```
Binary installer

Usage:
  binbrew [command]

Available Commands:
  help        Help about any command
  install     Install one or more binaries

Flags:
  -h, --help      help for binbrew
      --version   version for binbrew

Use "binbrew [command] --help" for more information about a command.
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
