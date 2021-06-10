##################################################################################
# VARIABLES
##################################################################################

variable "aws_access_key" {}
variable "aws_secret_key" {}
variable "vpc_name" {}
variable "private_key_path" {}
variable "key_name" {}
variable "database_password" {}
variable "database_port" {}
variable "region" {
  default = "us-east-1"
}

