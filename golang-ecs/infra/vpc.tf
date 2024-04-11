
# ----------------------
# VPC
# ----------------------
resource "aws_vpc" "golang-ecs-app-vpc" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true
  tags = {
    Name = "golang-ecs-app-vpc"
  }
}

# ----------------------
# サブネット
# ----------------------
resource "aws_subnet" "golang-ecs-app-subnet" {
  vpc_id                  = aws_vpc.golang-ecs-app-vpc.id
  cidr_block              = "10.0.1.0/24"
  availability_zone       = "ap-northeast-1a"
  map_public_ip_on_launch = true
  tags = {
    Name = "golang-ecs-app-subnet"
  }
}

# ----------------------
# セキュリティグループ
# ----------------------
resource "aws_security_group" "golang-ecs-app-sg" {
  name        = "golang-ecs-app-sg"
  description = "Allow all traffic"
  vpc_id      = aws_vpc.golang-ecs-app-vpc.id

  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "golang-ecs-app-sg"
  }
}
