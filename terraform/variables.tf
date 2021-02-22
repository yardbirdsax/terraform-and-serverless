variable aws_region_name {
  type = string
  description = "The region where AWS resources will be deployed."
  default = "us-east-2"
}

variable "deployment_name" {
  type = string
  description = "A unique name used to generate names for various resources. Should be unique at the Account level."
}