# Terrastate

[![CircleCI](https://circleci.com/gh/janritter/terrastate.svg?style=svg)](https://circleci.com/gh/janritter/terrastate)
[![Maintainability](https://api.codeclimate.com/v1/badges/235b50a37a1d73929d5c/maintainability)](https://codeclimate.com/github/janritter/terrastate/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/235b50a37a1d73929d5c/test_coverage)](https://codeclimate.com/github/janritter/terrastate/test_coverage)

> Tool to manage multiple states in Terraform - Allows Multi account setups

## Usage

### Generate statefile in the current directory

``` bash
terrastate --var-file ../../dev.tfvars
```

### Generate statefile, init and apply in the current directory

``` bash
terrastate apply --var-file ../../dev.tfvars
```

### Generate statefile, init and apply in the current directory with additional terraform flags

``` bash
terrastate apply --var-file ../../dev.tfvars -- -auto-approve
```

This generates a statefile called terrastate.tf

It will also replace an exisiting terrastate.tf file with the new information.

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

## Shell shorthand functions

Paste these functions into your ~/.zshrc or ~/.bashrc file and adapt the var-file path according to your project structure

ts was chosen as a shorthand for terrastate, you could also use tf instead 

```bash
function tsplan {
    echo terrastate plan --var-file ../test/$@.tfvars;
    terrastate plan --var-file ../test/$@.tfvars;
}

function tsaplly {
    echo terrastate apply --var-file ../test/$@.tfvars;
    terrastate apply --var-file ../test/$@.tfvars;
}

function tsdestroy {
    echo terrastate destroy --var-file ../test/$@.tfvars;
    terrastate destroy --var-file ../test/$@.tfvars;
}
```

## Additional Terrastate configuration

Terrastate offers the following options to set additional configuration

### Terrastate config file

The terrastate config file is a yaml file which can be either specified via `--terrastate-file` flag or is loaded from `$HOME/.terrastate.yaml`

#### Example content

```yaml
tf-init-upgrade: True
```

### Environment variable

Environemt variables prefixed with `TERRASTATE_` are automatically loaded by Terrastate, see the section for available variables below

### Flags

Some config options can be set via subcommand flag, e.g. `--tf-init-upgrade`

### List of available Terrastate config options

#### Upgrade for terraform init

Terraform init offers an `-upgrade` option which automatically updates all modules and plugins during init.
By specifying one of the following Terrastate config variables, this init option will be added by Terrastate.

The default is: `false`

- Config file:
  - Key is `tf-init-upgrade`
  - Value can be `True` or `False`
- Flag
  - The flag `--tf-init-upgrade` added to apply, plan, destroy or refresh commands
- Environment variable
  - `TERRASTATE_TF_INIT_UPGRADE=True` or `TERRASTATE_TF_INIT_UPGRADE=False`

## Installation

### For Mac using Homebrew

Get the formula
```
brew tap janritter/terrastate
```

Install terrastate
```
brew install terrastate
```

### For Linux

1. Download the latest release binary for your OS - [Releases](https://github.com/janritter/terrastate/releases) - For Linux (64bit) this would be 'linux_amd64_terrastate'
2. Rename the downloaded binary to terrastate and move it to your '/usr/bin/' or '/usr/local/bin' directory
3. Start using terrastate

## Development

### Resolve dependencies

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