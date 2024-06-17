resource "aws_subnet" "golang-study-public-a" {
  vpc_id                  = aws_vpc.golang-study-vpc.id
  cidr_block              = "10.0.1.0/24"
  availability_zone       = "ap-northeast-1a"
  map_public_ip_on_launch = true

  tags = {
    Name = "golang-study-public-a"
  }
}

resource "aws_subnet" "golang-study-private-a" {
  vpc_id                  = aws_vpc.golang-study-vpc.id
  cidr_block              = "10.0.2.0/24"
  availability_zone       = "ap-northeast-1a"
  map_public_ip_on_launch = false

  tags = {
    Name = "golang-study-private-a"
  }
}

resource "aws_subnet" "golang-study-public-c" {
  vpc_id                  = aws_vpc.golang-study-vpc.id
  cidr_block              = "10.0.3.0/24"
  availability_zone       = "ap-northeast-1c"
  map_public_ip_on_launch = true

  tags = {
    Name = "golang-study-public-c"
  }
}

resource "aws_subnet" "golang-study-private-c" {
  vpc_id                  = aws_vpc.golang-study-vpc.id
  cidr_block              = "10.0.4.0/24"
  availability_zone       = "ap-northeast-1c"
  map_public_ip_on_launch = false

  tags = {
    Name = "golang-study-private-c"
  }
}
