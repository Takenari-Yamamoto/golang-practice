# Terraform を使用して AWS RDS (Amazon Relational Database Service) のデータベースインスタンスを管理するためのリソース
resource "aws_db_instance" "golang-study-db" {
  # データベースインスタンスの一意の識別子です。AWS コンソール内でインスタンスを簡単に識別できるように
  identifier = "golang-study-rds"
  # 使用するデータベースの種類を指定します。ここでは postgres が使用されていますが、他にも mysql, oracle, sqlserver などが利用可能
  engine = "postgres"
  # データベースインスタンスのタイプ（サイズ）を指定します。ここでは db.t3.micro が使用されており、これは比較的小規模なインスタンス
  instance_class = "db.t3.micro"
  # データベースインスタンスに割り当てるストレージの量をギガバイト単位で指定
  allocated_storage = 20
  # データベースにアクセスするためのユーザ名とパスワードです。これらはデータベースへのログインに使用される
  username = "dbuser"
  password = "dbpassword"
  #  データベースの名前を指定します。この名前は、特定のデータベースエンジンの要件に従う必要がある
  name = "golang_study_db"
  # データベースインスタンスに適用するセキュリティグループのIDのリストです。これにより、インスタンスへのネットワークアクセスを制御
  vpc_security_group_ids = [
    aws_security_group.golang-study-db-sg.id
  ]
  # インスタンスが配置されるサブネットグループを指定します。これは、データベースが特定のVPC内の特定のサブネットに配置されることを保証
  db_subnet_group_name = aws_db_subnet_group.golang-study-db_subnet_group.name
  # ルチアベイラビリティーゾーンの高可用性オプションを有効にするかどうかを指定します。
  # これが true に設定されると、プライマリデータベースのスタンバイコピーが異なるアベイラビリティーゾーンに作成され、
  # 障害発生時に自動的にフェイルオーバーが行われます
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
