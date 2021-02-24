locals {
  db_admin_user_name = "pgadmin"
  db_app_user_name   = "app"
}

resource "random_password" "rds" {
  length = 12
  special = false
}

resource "aws_security_group" "rds" {
  name = "${var.deployment_name}-rds"
}

resource "aws_security_group_rule" "rds-ingress" {
  security_group_id = aws_security_group.rds.id
  cidr_blocks = [ "0.0.0.0/0" ] # This is a REALLY bad idea for anything other than a sandbox
  from_port = 5432
  to_port = 5432
  protocol = "tcp"
  type = "ingress"
}

resource "aws_security_group_rule" "rds-egress" {
  security_group_id = aws_security_group.rds.id
  cidr_blocks = [ "0.0.0.0/0" ]
  from_port = 0
  to_port = 0
  protocol = "-1"
  type = "egress"
}

resource "aws_db_instance" "rds" {
  name = replace(var.deployment_name,"-","")
  engine = "postgres"
  allocated_storage = 20
  instance_class = "db.t3.micro"
  username = local.db_admin_user_name
  password = random_password.rds.result
  skip_final_snapshot = true
  tags = {
    "deployment" = var.deployment_name
  }
  publicly_accessible = true
  vpc_security_group_ids = [ aws_security_group.rds.id ]
}

output rds_endpoint {
  value = aws_db_instance.rds.endpoint
}