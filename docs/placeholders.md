# Placeholders

Terrastate supports placeholders in all string variables.

## How to use?

Placeholders are included in the variable value.

Example: my-project/{{ current.dir }}/ec2

## Supported placeholders

### current.dir

This variable cointains the name of the directory terrastate is currently executed in.

:warning: Since there are multiple keys for multiple Terraform subdirectories required, the `state_key` variable must contain the placeholder {{ current.dir }}.
