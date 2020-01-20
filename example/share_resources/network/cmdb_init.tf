# create dynamoDB item
resource "aws_dynamodb_table_item" "init" {
  provider = aws.audit
  hash_key = var.share_resources_hash_key
  table_name = var.share_resources_dynamodb
  item = <<ITEM
{
  "${var.share_resources_hash_key}": {"S": "${var.aws}.${var.prefix}.vpc.default"},
  "vpc_id": {"S": "${aws_vpc.default.id}"},
  "vpc_cidr": {"S": "${aws_vpc.default.cidr_block}"},
  "subnet_1_pub": {"S": "${aws_subnet.subnet_1_pub.id}"},
  "subnet_2_pub": {"S": "${aws_subnet.subnet_2_pub.id}"},
  "subnet_1_pub_cidr": {"S": "${aws_subnet.subnet_1_pub.cidr_block}"},
  "subnet_2_pub_cidr": {"S": "${aws_subnet.subnet_2_pub.cidr_block}"},
  "subnet_3_pub": {"S": "${aws_subnet.subnet_3_pub.id}"},
  "subnet_3_pub_cidr": {"S": "${aws_subnet.subnet_3_pub.cidr_block}"}
}
ITEM

}