# create dynamoDB item
resource "aws_dynamodb_table_item" "init" {
  provider = aws.audit
  hash_key = var.share_resources_hash_key
  table_name = var.share_resources_dynamodb
  item = <<ITEM
{
  "${var.share_resources_hash_key}": {"S": "${var.aws}.${var.prefix}.${var.service_name}"},
  "creation_delay": {"S": "${var.creation_delay}"}

}
ITEM

}