resource "aws_internet_gateway" "golang-study-igw" {
  vpc_id = aws_vpc.golang-study-vpc.id

  tags = {
    Name = "golang-study-igw"
  }
}
