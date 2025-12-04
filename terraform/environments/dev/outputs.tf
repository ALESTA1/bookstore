# outputs.tf

output "vpc_id" {
  description = "ID of the VPC"
  value       = aws_vpc.dev_vpc.id
}

output "vpc_cidr_block" {
  description = "CIDR block of the VPC"
  value       = aws_vpc.dev_vpc.cidr_block
}

output "internet_gateway_id" {
  description = "ID of the Internet Gateway"
  value       = aws_internet_gateway.dev_igw.id
}

output "subnet_id" {
  description = "ID of the subnet"
  value       = aws_subnet.dev_subnet.id
}

output "subnet_cidr_block" {
  description = "CIDR block of the subnet"
  value       = aws_subnet.dev_subnet.cidr_block
}

output "subnet_availability_zone" {
  description = "Availability zone of the subnet"
  value       = aws_subnet.dev_subnet.availability_zone
}

output "route_table_id" {
  description = "ID of the route table"
  value       = aws_route_table.dev_rt.id
}

output "security_group_id" {
  description = "ID of the security group"
  value       = aws_security_group.dev_sg.id
}

output "security_group_name" {
  description = "Name of the security group"
  value       = aws_security_group.dev_sg.name
}

output "key_pair_name" {
  description = "Name of the EC2 key pair"
  value       = aws_key_pair.dev_key.key_name
}

output "instance_id" {
  description = "ID of the EC2 instance"
  value       = aws_instance.dev_instance.id
}

output "instance_public_ip" {
  description = "Public IP address of the EC2 instance"
  value       = aws_instance.dev_instance.public_ip
}

output "instance_private_ip" {
  description = "Private IP address of the EC2 instance"
  value       = aws_instance.dev_instance.private_ip
}

output "instance_public_dns" {
  description = "Public DNS name of the EC2 instance"
  value       = aws_instance.dev_instance.public_dns
}

output "instance_state" {
  description = "State of the EC2 instance"
  value       = aws_instance.dev_instance.instance_state
}

output "instance_type" {
  description = "Type of the EC2 instance"
  value       = aws_instance.dev_instance.instance_type
}

output "ssh_connection_command" {
  description = "SSH command to connect to the instance"
  value       = "ssh -i ~/.ssh/bookstore-dev-key ec2-user@${aws_instance.dev_instance.public_ip}"
}

output "deployment_info" {
  description = "Deployment information"
  value = {
    environment         = var.environment
    project            = var.project_name
    region             = var.region
    deployment_time    = timestamp()
    total_resources    = 8
  }
}
