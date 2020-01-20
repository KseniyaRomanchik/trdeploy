terraform {
  backend "s3" {}
}

# aws provider config
provider "aws" {
  alias = "audit"
  region                  = var.region
  profile                 = var.aws_audit
}

provider "aws" {
  alias = "work"
  region                  = var.region
  profile                 = var.aws
}
