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

provider "postgresql" {
  host = aws_db_instance.rds.address
  port = aws_db_instance.rds.port
  database = aws_db_instance.rds.name
  username = local.db_admin_user_name
  password = random_password.rds.result
  superuser = false
}