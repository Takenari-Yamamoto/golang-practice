resource "aws_route_table" "golang-study-public" {
  vpc_id = aws_vpc.golang-study-vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.golang-study-igw.id
  }

  tags = {
    Name = "golang-study-public"
  }
}

resource "aws_route_table_association" "golang-study-public-association" {
  subnet_id      = aws_subnet.golang-study-public-a.id
  route_table_id = aws_route_table.golang-study-public.id
}
