# ----------------------
# インターネットゲートウェイの設定
# ----------------------
resource "aws_internet_gateway" "golang_ecs_app_igw" {
  vpc_id = aws_vpc.golang-ecs-app-vpc.id

  tags = {
    Name = "golang-ecs-app-igw"
  }
}

# ----------------------
# ルートテーブルの設定
# ----------------------
resource "aws_route_table" "golang_ecs_app_rt" {
  vpc_id = aws_vpc.golang-ecs-app-vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.golang_ecs_app_igw.id
  }

  tags = {
    Name = "golang-ecs-app-rt"
  }
}

resource "aws_route_table_association" "golang_ecs_app_rta" {
  subnet_id      = aws_subnet.golang-ecs-app-subnet.id
  route_table_id = aws_route_table.golang_ecs_app_rt.id
}
