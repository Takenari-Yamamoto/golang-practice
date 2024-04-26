resource "aws_lb" "golang-study-api-public" {
  name = "golang-study-api-public"
  # ロードバランサーが内部向けか外部向けかを示します。false はインターネット向けです。
  internal = false
  #  ロードバランサーのタイプを指定します。application はアプリケーション層でのロードバランシングを意味します。
  load_balancer_type = "application"
  # ALBが属するセキュリティグループ。トラフィックのフィルタリングを行います。
  security_groups = [aws_security_group.golang-study-alb-sg.id]
  # ALBが配置されるサブネット。複数のサブネットを指定することで高可用性が保たれます
  subnets = [
    aws_subnet.golang-study-public-a.id,
    aws_subnet.golang-study-public-c.id
  ]

  enable_deletion_protection = false

  tags = {
    Name = "golang-study-api-public"
  }
}

# ALBによってトラフィックがルーティングされるEC2インスタンスやコンテナなどのグループ
resource "aws_lb_target_group" "golang-study-api-public-tg" {
  name = "golang-study-api-public-tg"
  #  バックエンドと通信する際のポートとプロトコル
  port        = 80
  protocol    = "HTTP"
  vpc_id      = aws_vpc.golang-study-vpc.id
  target_type = "ip"

  # ターゲットの健康状態をチェックするための設定です。正常に動作しているかを確認し、問題があるインスタンスからはトラフィックを引き離します
  health_check {
    enabled             = true
    interval            = 30
    path                = "/"
    port                = "traffic-port"
    protocol            = "HTTP"
    healthy_threshold   = 2
    unhealthy_threshold = 3
    timeout             = 5
    matcher             = "200"
  }

  tags = {
    Name = "golang-study-api-public-tg"
  }

}

# リスナーは、ALBがリクエストを受け取る際のエントリーポイントです。
resource "aws_lb_listener" "golang-study-api-public-listener" {
  load_balancer_arn = aws_lb.golang-study-api-public.arn
  # リスナーが受け取るトラフィックのポートとプロトコル。
  port     = "80"
  protocol = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.golang-study-api-public-tg.arn
  }
}
