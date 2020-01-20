resource "aws_subnet" "subnet_1_pub" {
  provider = aws.work
  vpc_id     = aws_vpc.default.id
  cidr_block = var.subnet_1_pub_default
  availability_zone = data.aws_availability_zones.available.names[0]
  tags = {
    Name = "${var.aws}.${var.prefix}.default.subnet_1_pub"
  }
  lifecycle  {
    ignore_changes = [tags]
  }
}

resource "aws_subnet" "subnet_2_pub" {
  provider = aws.work
  vpc_id     = aws_vpc.default.id
  cidr_block = var.subnet_2_pub_default
  availability_zone = data.aws_availability_zones.available.names[1]
  tags = {
    Name = "${var.aws}.${var.prefix}.default.subnet_2_pub"
  }
    lifecycle  {
    ignore_changes = [tags]
  }
}

resource "aws_subnet" "subnet_3_pub" {
  provider = aws.work
  vpc_id     = aws_vpc.default.id
  cidr_block = var.subnet_3_pub_default
  availability_zone = data.aws_availability_zones.available.names[2]
  tags  = {
    Name = "${var.aws}.${var.prefix}.default.subnet_3_pub"
  }
    lifecycle  {
    ignore_changes = [tags]
  }
}


