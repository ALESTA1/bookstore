# main.tf

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.0"
    }
  }
}

provider "aws" {
  region = var.region
}

module "vpc" {
  source = "terraform-aws-modules/vpc/aws"
  version = ">= 3.0"

  name = "${var.project_name}-${var.environment}-vpc"
  cidr = var.vpc_cidr

  azs                 = ["${var.region}a", "${var.region}b"]
  private_subnets = var.private_subnet_cidrs
  public_subnets  = var.public_subnet_cidrs

  enable_nat_gateway = true
  single_nat_gateway = true
  enable_vpn_gateway = false

  tags = var.tags
}

resource "aws_security_group" "allow_ssh" {
  name          = "${var.project_name}-${var.environment}-allow-ssh"
  description = "Allow SSH inbound traffic"
  vpc_id      = module.vpc.vpc_id

  ingress {
    description = "SSH from VPC"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = [var.vpc_cidr]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = var.tags
}

resource "aws_instance" "dev_instance" {
  ami               = var.ami_id
  instance_type = var.instance_type
  key_name        = var.key_name

  subnet_id                              = module.vpc.private_subnets[0]
  vpc_security_group_ids               = [aws_security_group.allow_ssh.id]
  associate_public_ip_address = false

  tags = merge(var.tags, {
    Name = "${var.project_name}-${var.environment}-instance"
  })
}
