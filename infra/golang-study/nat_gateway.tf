# Elastic IPは、AWSの静的なパブリックIPアドレスです。NATゲートウェイにはElastic IPが必要で、
# これによってプライベートサブネットのインスタンスがインターネットと通信できるようになります。
resource "aws_eip" "golang-study-eip" {
  vpc = true
}

# NATゲートウェイは、プライベートサブネットのインスタンスがインターネットへのアウトバウンド接続を行うためのサービスです。
# インバウンド接続は許可されません。
resource "aws_nat_gateway" "golang-study-ngw" {
  allocation_id = aws_eip.golang-study-eip.id
  subnet_id     = aws_subnet.golang-study-public-a.id

  tags = {
    Name = "golang-study-ngw"
  }
}
