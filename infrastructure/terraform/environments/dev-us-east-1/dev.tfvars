# dev.tfvars

region                   = "us-east-1"
project_name              = "dev-project"
environment               = "dev"
vpc_cidr                  = "10.0.0.0/16"
private_subnet_cidrs     = ["10.0.1.0/24", "10.0.2.0/24"]
public_subnet_cidrs      = ["10.0.101.0/24", "10.0.102.0/24"]
instance_type             = "t3.micro"
ami_id                    = "ami-0c02fb55956c7d316"
key_name                  = "your-key-pair-name"

tags = {
  Environment = "dev"
  Project     = "dev-project"
  ManagedBy   = "terraform"
  CostCenter  = "development"
}
