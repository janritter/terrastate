# Terrastate - S3 Backend

:warning: Currently there are only some of the Terraform supported variables for the S3 backend implemented.

## Variables

All supported variables can be found in the `backendAttributes` struct in the [backend/s3/main.go](backend/s3/main.go) file.

The `VarKey` is the variable key you should use in the variable file.
The `AttributeKey` will be used as attribute key in the generated terraform state file.

### Additional information about placeholders

[Placeholder README](docs/placeholder.md)

