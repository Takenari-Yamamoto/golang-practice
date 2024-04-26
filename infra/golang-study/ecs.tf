resource "aws_ecr_repository" "golang-study-api-repo" {
  name                 = "golang-study-api-repo"
  image_tag_mutability = "MUTABLE"
}
