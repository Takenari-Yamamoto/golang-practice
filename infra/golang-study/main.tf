terraform {
  required_version = ">= 0.13"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~>3.0"
    }
  }
  backend "s3" {
    bucket  = "golang-study-terraform-state"
    key     = "appserver/terraform.tfstate"
    region  = "ap-northeast-1"
    profile = "terraform_user"
  }

  # backend "local" {}
}

# ----------------------
# provider configuration
# ----------------------
provider "aws" {
  profile = "terraform_user"
  region  = "ap-northeast-1"
}
