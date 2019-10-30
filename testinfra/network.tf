resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"

  tags = {
    Name = "terraform-provider-ciscoasa acc test"
  }
}

resource "aws_internet_gateway" "gw" {
  vpc_id = "${aws_vpc.main.id}"
}

resource "aws_subnet" "main" {
  availability_zone = "${data.aws_availability_zones.available.names[0]}"
  cidr_block        = "10.0.0.0/24"
  vpc_id            = "${aws_vpc.main.id}"

  tags = {
    Name = "terraform-provider-ciscoasa acc test"
  }
}

resource "aws_route_table" "main" {
  vpc_id = "${aws_vpc.main.id}"
}

resource "aws_route" "public_internet_gateway" {
  route_table_id         = "${aws_route_table.main.id}"
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = "${aws_internet_gateway.gw.id}"
}

resource "aws_route_table_association" "main" {
  route_table_id = "${aws_route_table.main.id}"
  subnet_id      = "${aws_subnet.main.id}"
}
