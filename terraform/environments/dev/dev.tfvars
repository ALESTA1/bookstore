# dev.tfvars

# AWS Configuration
region        = "us-east-1"
vpc_cidr      = "10.0.0.0/16"
subnet_cidr   = "10.0.1.0/24"

# EC2 Configuration
instance_type     = "t3.micro"
ami_id            = "ami-0c02fb5595656c7d316"  # Amazon Linux 2023 AMI (HVM) - us-east-1
root_volume_size = 20

# Security Configuration
allowed_ssh_cidrs = [
  "0.0.0.0/0"  # CHANGE THIS: Replace with your specific IP range for better security
  # Example: ["203.0.113.0/24", "198.51.100.0/24"]
]

# SSH Key Configuration
public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC7..."  # REPLACE WITH YOUR ACTUAL PUBLIC KEY

# Project Tags
environment  = "dev"
project_name = "bookstore"
owner        = "raghav.arora@unifyapps.com"
cost_center  = "development"
