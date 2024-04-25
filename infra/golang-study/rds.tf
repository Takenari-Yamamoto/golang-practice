resource "aws_db_instance" "golang-study-db" {
  identifier        = "golang-study-rds"
  engine            = "postgres"
  instance_class    = "db.t3.micro"
  allocated_storage = 20
  username          = "dbuser"
  password          = "dbpassword"
  name              = "golang_study_db"

  vpc_security_group_ids = [
    aws_security_group.golang-study-db-sg.id
  ]
  db_subnet_group_name = aws_db_subnet_group.golang-study-db_subnet_group.name

  multi_az = false

  tags = {
    Name = "golang-study-rds"
  }
}

# RDSインスタンスをデプロイするために使用するサブネットのグループ
# このリソースは、特定のVPC内の複数のサブネットをグループ化し、RDSインスタンスに割り当てます
resource "aws_db_subnet_group" "golang-study-db_subnet_group" {
  name       = "golang-study-db-subnet-group"
  subnet_ids = [aws_subnet.golang-study-private-a.id, aws_subnet.golang-study-private-c.id]
}
