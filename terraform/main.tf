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

resource "random_password" "rds" {
  length = 12
  special = false
}

resource "aws_db_instance" "rds" {
  name = replace(var.deployment_name,"-","")
  engine = "postgres"
  allocated_storage = 20
  instance_class = "db.t3.micro"
  username = "pgadmin"
  password = random_password.rds.result
  skip_final_snapshot = true
  tags = {
    "deployment" = var.deployment_name
  }
  publicly_accessible = true
}

output rds_endpoint {
  value = aws_db_instance.rds.endpoint
}