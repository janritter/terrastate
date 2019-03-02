# Terrastate

[![Build Status](https://travis-ci.com/janritter/terrastate.svg?token=fPhMwJC3SnTkQrfzte44&branch=master)](https://travis-ci.com/janritter/terrastate)
[![Maintainability](https://api.codeclimate.com/v1/badges/235b50a37a1d73929d5c/maintainability)](https://codeclimate.com/github/janritter/terrastate/maintainability)

> Tool to manage multiple states in Terraform - Allows Multi account setups

## Use case

> TODO

## Usage

### Generate statefile in the current directory

``` bash
terrastate --var-file ../../dev.tfvars
```

This generates a statefile called terrastate.tf

It will also replace an exisiting terrastate.tf file with the new infromation.

#### Required Terraform variables

Some variables must be set in the varfile used by terrastate. Some of these variables are used by terrastate and others are depending on the backend.

The state_backend var defines the backend type and therefore which backend implementation should be used by terrastate. If the given backend is not supported terrastate returns an error.
In this example s3 is configured as backend.

```bash
state_backend = "s3"
```

The currently supported backends are:

- [S3](docs/s3-backend.md)

Feel free to add a new backend and create a pull request - [How to create a new backend?](docs/own-backend.md)

For more information about the backend specific variables click the backend in the list above.

### Version

``` bash
terrastate version
```

### Help

``` bash
terrastate help
```

## Development

### Resolve dependencies

Requires go dep

```make
make prepare
```

### Build binary

compiles to bin/terrastate

```make
make build
```

## License and Author

Author: Jan Ritter

License: MIT