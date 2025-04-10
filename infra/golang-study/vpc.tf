resource "aws_vpc" "golang-study-vpc" {
  cidr_block = "10.0.0.0/16"
  tags = {
    Name = "golang-study-vpc"
  }
}
