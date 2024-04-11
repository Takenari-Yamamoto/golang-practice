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
