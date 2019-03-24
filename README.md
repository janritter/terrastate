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

#### Optional Terraform variables

##### Automatic removal of the .terraform folder

```bash
state_auto_remove_old = true
```

When you set the value to true, the .terraform folder in the current directory gets removed when you create a terrastate.tf backend configuration through terrastate. If you set the value to false or don't set it, then the .terraform folder will not be removed.

This option allowes you to execute terrastate and then directly terraform init without manually removing the .terraform folder.

### Version

``` bash
terrastate version
```

### Help

``` bash
terrastate help
```

## Installation

1. Download the latest release binary for your OS - [Releases](https://github.com/janritter/terrastate/releases) - For Mac (64bit) this would be 'darwin_amd64_terrastate'
2. Rename the downloaded binary to terrastate and move it to your '/usr/bin/' or '/usr/local/bin' directory
3. Start using terrastate

## Development

### Resolve dependencies

"go dep" must be installed

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