
# ----------------------
# ルートテーブルの作成
# パブリックサブネットがインターネットゲートウェイに接続できるように設定
# ----------------------
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

resource "aws_route_table_association" "golang-study-public-a-association" {
  subnet_id      = aws_subnet.golang-study-public-a.id
  route_table_id = aws_route_table.golang-study-public.id
}

# --------------------　ここからプライベートサブネットのルートテーブルを設定 ---------------
# プライベートサブネットのルートテーブルを設定して、
# 全てのインターネット向けトラフィック (0.0.0.0/0) をNATゲートウェイへルーティングします。
# これにより、プライベートサブネットのインスタンスがインターネットにアクセスできるようになります。
# ---------------------------------------------------------------------------------
resource "aws_route_table" "golang-study-private" {
  vpc_id = aws_vpc.golang-study-vpc.id

  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.golang-study-ngw.id
  }

  tags = {
    Name = "golang-study-private"
  }
}

resource "aws_route_table_association" "golang-study-private-a-association" {
  subnet_id      = aws_subnet.golang-study-private-a.id
  route_table_id = aws_route_table.golang-study-private.id
}
