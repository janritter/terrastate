terraform {
  backend "s3" {
    region = "eu-central-1"
    bucket = "local-test-bucket"
    key    = "terrastate/terrastate-e2e/terraform.tfstate"
    acl    = "bucket-owner-full-control"
  }
}