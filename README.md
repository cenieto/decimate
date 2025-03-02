# Decimate

[![Go Reference](https://pkg.go.dev/badge/github.com/cenieto/decimator.svg)](https://pkg.go.dev/github.com/cenieto/decimator)
[![Go Version](https://img.shields.io/github/go-mod/go-version/cenieto/decimator)](https://github.com/cenieto/decimator)
[![Go Report Card](https://goreportcard.com/badge/github.com/cenieto/decimator)](https://goreportcard.com/report/github.com/cenieto/decimator)
[![Go Reference](https://pkg.go.dev/badge/github.com/cenieto/decimate.svg)](https://pkg.go.dev/github.com/cenieto/decimate)

## Overview

Decimate is a Go package designed to reduce datasets while preserving essential information. Currently, it supports operations on 2D and 3D linestrings using Euclidean distance.

To enable decimation algorithms, several geometric operations are included. The `gonum` package is utilized to facilitate these geometric operations and matrix computations.

## Dependencies

decimate depends on gonum, ensure you have gonum installed:
```go
go get gonum.org/v1/gonum/mat
```

## Installation

To install decimate, ensure you have Go installed and run:

```sh
go get github.com/cenieto/decimator@latest
```

To import and use it in your project:

```go
import "github.com/cenieto/decimator"
```

## Usage Example

Currently, no decimation algorithm is implemented, but the example below demonstrates a 2D operation on vectors. Additional examples can be found in the test implementations.

The cross product of two 2D vectors can be computed as follows:

TODO: the example is incorrect. Package name is not decimator animore.

```go
package main

import (
	"fmt"
	"decimator/pkg/primitives"
	"decimator/pkg/geom2d"
)

func main() {
	v1 := primitives.NewVector([]float64{1, 2})
	v2 := primitives.NewVector([]float64{3, 4})
	geom := geom2d.NewEuclid()
	result := geom.CrossProduct(v1, v2)
	fmt.Println(result)
}
```

## Stack Script

The `stack` script is a Bash script that automates common tasks. Before using it, grant execution permissions:

```sh
chmod +x stack
```

### Features

The `stack` script provides the following functionalities:

- Builds a lightweight container to execute `decimate` functionalities or run tests in a reproducible and consistent environment.
- Deploys a container with the current folder mounted to `/app` inside the container.
- Provides interactive access to the containerized environment via a Bash shell.
- Runs code inside the container.
- Executes unit tests.
- Formats all `.go` files according to Go standards.
- Runs linting to ensure code quality.
- Installs necessary dependencies for the above workflows.

To see available options, run:

```sh
./stack help
```

## CI/CD Workflow

This repository includes a CI/CD pipeline using Docker containers and the stack script to automate testing.

### Unit Testing Workflow Steps

1. **Build:** Creates a Docker container to ensure a consistent environment.
2. **Deploy Container:** Launches the container with the current folder mounted to `/app`.
3. **Run Tests:** Executes all unit tests inside the container.

```sh
./stack build
./stack up 
./stack unit ./...
```

### Linting & Static Analysis

The project uses `golangci-lint` to maintain code quality and enforce best practices.

These workflows are also executed automatically in a GitHub Actions environment using [`.github/workflows/test.yaml`](.github/workflows/test.yaml).

## License

This project is licensed under the Apache 2.0 License. See [LICENSE](LICENSE) for details.


