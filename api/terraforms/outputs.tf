output "ecs_service_dns_name" {
  value = aws_ecs_service.url-encoder_service.network_configuration[0].subnets
}

output "ecr_repository_url" {
  description = "The URI of the created ECR repository"
  value       = aws_ecr_repository.url_encoder.repository_url
}
