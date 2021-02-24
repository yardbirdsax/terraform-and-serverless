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