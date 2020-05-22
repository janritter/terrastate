# Terrastate

state_backend = "s3"

state_bucket = "local-test-bucket"

state_dynamodb_table = "terraform-state-lock"

region = "eu-central-1"

state_key = "terrastate/{{ current.dir }}/terraform.tfstate"

state_auto_remove_old = true

stage = "test"
