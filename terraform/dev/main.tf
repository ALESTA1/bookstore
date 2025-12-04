# main.tf

provider "aws" {
  region = var.region
}

resource "aws_vpc" "dev_vpc" {
  cidr_block = var.vpc_cidr
  enable_dns_hostnames = true
  enable_dns_support = true

  tags = {
    Name = "dev-vpc"
    Environment = "dev"
  }
}

resource "aws_subnet" "dev_subnet" {
  vpc_id = aws_vpc.dev_vpc.id
  cidr_block = var.subnet_cidr
  availability_zone = "${var.region}a"

  tags = {
    Name = "dev-subnet"
    Environment = "dev"
  }
}

resource "aws_security_group" "dev_sg" {
  name = "dev-security-group"
  description = "Allow inbound traffic for development"
  vpc_id = aws_vpc.dev_vpc.id

  ingress {
    description = "SSH from anywhere"
    from_port = 22
    to_port = 22
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "dev-security-group"
    Environment = "dev"
  }
}

resource "aws_instance" "dev_instance" {
  ami = var.ami_id
  instance_type = var.instance_type
  subnet_id = aws_subnet.dev_subnet.id
  vpc_security_group_ids = [aws_security_group.dev_sg.id]

  tags = {
    Name = "dev-instance"
    Environment = "dev"
  }
}
