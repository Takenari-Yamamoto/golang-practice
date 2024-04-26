resource "aws_ecr_repository" "golang-study-api-repo" {
  name                 = "golang-study-api-repo"
  image_tag_mutability = "MUTABLE"
}

resource "aws_ecs_cluster" "golang_study_cluster" {
  name = "golang-study-cluster"
}
