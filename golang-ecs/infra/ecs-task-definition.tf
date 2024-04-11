# ----------------------
# ECSタスクの定義
# ----------------------
resource "aws_ecs_task_definition" "golang-ecs-app" {
  family                   = "golang-ecs-app"
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = "256"
  memory                   = "512"
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn

  container_definitions = jsonencode([
    {
      name      = "golang-ecs-app"
      image     = "${aws_ecr_repository.go-ecs-app.repository_url}:latest"
      essential = true
      portMappings = [
        {
          containerPort = 8080
          hostPort      = 8080
          protocol      = "tcp"
        }
      ]
    }
  ])
}

# ECSサービス
resource "aws_ecs_service" "golang_ecs_app_service" {
  name            = "golang-ecs-app-service"
  cluster         = aws_ecs_cluster.golang_ecs_cluster.id
  task_definition = aws_ecs_task_definition.golang-ecs-app.arn
  desired_count   = 1
  launch_type     = "FARGATE"

  network_configuration {
    subnets = [
      aws_subnet.golang-ecs-app-subnet.id
    ]
    security_groups = [
      aws_security_group.golang-ecs-app-sg.id
    ]
    assign_public_ip = true
  }

  depends_on = [
    aws_ecs_task_definition.golang-ecs-app
  ]
}
