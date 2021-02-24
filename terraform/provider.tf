terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~>3.26"
    }
    postgresql = {
      source = "cyrilgdn/postgresql"
      version = "1.11.1"
    }
  }
  required_version = ">0.13"
}

provider aws {
  region = var.aws_region_name
}