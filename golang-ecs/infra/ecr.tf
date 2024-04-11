resource "aws_ecr_repository" "go-ecs-app" {
  name                 = "go-ecs-app"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}
