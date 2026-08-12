# Terrastate - Parent Directory Placeholder Test

state_backend = "s3"

state_bucket = "local-test-bucket"

state_dynamodb_table = "terraform-state-lock"

region = "eu-central-1"

state_key = "terrastate/{{ parent.dir }}/terraform.tfstate"

state_auto_remove_old = true

state_acl = "bucket-owner-full-control"

stage = "test"
