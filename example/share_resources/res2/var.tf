variable "region" {}
variable "aws"{}
variable "aws_audit" {}
variable "env_dynamodb" {}
variable "app_dynamodb" {}
variable "share_resources_dynamodb" {}
variable "prefix" {}
variable "create_user" {
  default = "terraform" # it is from ci/cd
}
#common
variable "env_hash_key" {}
variable "share_resources_hash_key" {}

# module
variable "service_name" {}
variable "creation_delay" {}
variable "destroy_delay" {}