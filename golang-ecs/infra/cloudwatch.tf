resource "aws_cloudwatch_log_group" "ecs_log_group" {
  name = "/ecs/golang-ecs-app"
}
