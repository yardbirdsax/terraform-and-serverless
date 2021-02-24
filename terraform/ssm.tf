resource "aws_ssm_parameter" "rds_connection_string" {
  name = "/env/app/rds-connection-string"
  type = "SecureString"
  value = "host=${aws_db_instance.rds.address} user=${local.db_app_user_name} password=${random_password.app.result} DB.name=${aws_db_instance.rds.name} port=${aws_db_instance.rds.port} sslmode=enable"
}