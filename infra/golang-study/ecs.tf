resource "aws_ecr_repository" "golang-study-api-repo" {
  name                 = "golang-study-api-repo"
  image_tag_mutability = "MUTABLE"
}

resource "aws_ecs_cluster" "golang_study_cluster" {
  name = "golang-study-cluster"
}

/*
** ECSのタスク定義
** ECRからプッシュしたイメージの使用、必要なCPUとメモリの指定などを行う
*/
resource "aws_ecs_task_definition" "golang_study_app" {
  family                   = "golang-study-app"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "256"
  memory                   = "512"
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn

  container_definitions = jsonencode([
    {
      name      = "golang-study-app"
      image     = "${aws_ecr_repository.golang-study-api-repo.repository_url}:latest"
      cpu       = 256
      memory    = 512
      essential = true
      portMappings = [
        {
          containerPort = 80
          hostPort      = 80
          protocol      = "tcp"
        }
      ],

      // TODO: secreat manager 的なやつで管理したい
      environment = [
        {
          name  = "DATABASE_HOST"
          value = aws_db_instance.golang-study-db.address
        },
        {
          name  = "DATABASE_PORT"
          value = "5432"
        },
        {
          name  = "DATABASE_USER"
          value = "dbuser"
        },
        {
          name  = "DATABASE_PASSWORD"
          value = "dbpassword"
        },
        {
          name  = "DATABASE_NAME"
          value = "golang_study_db"
        }
      ]
    }
  ])
}

/* 
** ECSのタスクを実行するためのIAMロール
** AWSのサービスがECSタスクを適切に管理できるようにするために必要
*/
resource "aws_iam_role" "ecs_task_execution_role" {
  name = "ecs_task_execution_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        },
        Effect = "Allow",
      },
    ],
  })
}

resource "aws_iam_role_policy_attachment" "ecs_task_execution_role_policy" {
  role       = aws_iam_role.ecs_task_execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}
