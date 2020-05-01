resource "aws_vpc" "default" {
  provider = aws.work
  cidr_block       = var.vpc_default_cidr
  enable_dns_support=true
  enable_dns_hostnames=true
  tags = {
    Name = "${var.aws}.${var.prefix}.default"
  }
    lifecycle   {
    ignore_changes = [tags]
  }
}

resource "aws_internet_gateway" "default" {
  provider = aws.work
  vpc_id = aws_vpc.default.id
   tags = {
    Name = "${var.aws}.${var.prefix}.default"
  }
    lifecycle {
    ignore_changes = [tags]
  }
}


