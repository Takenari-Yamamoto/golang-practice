resource "aws_cloudwatch_log_group" "ecs_logs" {
  name              = "/ecs/golang-study-app"
  retention_in_days = 30
}
