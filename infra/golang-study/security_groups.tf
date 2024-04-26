resource "aws_security_group" "golang-study-sg" {
  vpc_id = aws_vpc.golang-study-vpc.id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 443
    to_port     = 443
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
    Name = "golang-study-sg"
  }
}

resource "aws_security_group" "golang-study-alb-sg" {
  vpc_id = aws_vpc.golang-study-vpc.id

  ingress {
    from_port   = 80
    to_port     = 80
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
    Name = "golang-study-alb-sg"
  }
}

resource "aws_security_group" "golang-study-db-sg" {

  name        = "golang-study-db-sg"
  description = "security group for RDS"
  vpc_id      = aws_vpc.golang-study-vpc.id

  // PostgreSQLのデフォルトポートである5432番ポートのトラフィックのみを許可
  ingress {
    from_port = 5432
    to_port   = 5432
    protocol  = "tcp"
    security_groups = [
      aws_security_group.golang-study-sg.id
    ]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "golang-study-db-sg"
  }

}
