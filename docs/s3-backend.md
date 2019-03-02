# Terrastate - S3 Backend

:warning: Currently there are only some of the Terraform supported variables for the S3 backend implemented.

## Variables

| Name in .tfvars      | Name in Terraform backend config | Example                                        |
| -------------------- | :------------------------------: | ---------------------------------------------- |
| state_bucket         |              bucket              | my-terraform-state                             |
| state_dynamodb_table |          dynamodb_table          | terraform-state-lock                           |
| region               |              region              | eu-central-1                                   |
| state_key            |               key                | terrastate/{{ current.dir }}/terraform.tfstate |

### Additional information about the state_key

Since there are multiple keys for multiple Terraform subdirectories required, the key value contains the placeholder {{ current.dir }}.

{{ current.dir }} gets automaticly replaced with the current directory before writing the terrastate.tf file.

#### Examples

- terrastate gets executed in my-project/ec2/ this creates the key terrastate/ec2/terraform.tfstate
- terrastate gets executed in my-project/alb/ this creates the key terrastate/alb/terraform.tfstate
