resource "aws_route_table" "pub" {
  provider = aws.work
  vpc_id  = aws_vpc.default.id
  route  {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.default.id
  }

  tags = {
    Name = "${var.aws}.${var.prefix}.default.pub"
  }
    lifecycle   {
    ignore_changes = [tags]
  }
}


resource "aws_route_table_association" "pub_1" {
  provider = aws.work
  subnet_id      = aws_subnet.subnet_1_pub.id
  route_table_id = aws_route_table.pub.id
}

resource "aws_route_table_association" "pub_2" {
  provider = aws.work
  subnet_id      = aws_subnet.subnet_2_pub.id
  route_table_id = aws_route_table.pub.id
}

resource "aws_route_table_association" "pub_3" {
  provider = aws.work
  subnet_id      = aws_subnet.subnet_3_pub.id
  route_table_id = aws_route_table.pub.id
}
