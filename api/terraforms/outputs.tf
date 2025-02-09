output "ecs_service_dns_name" {
  value = aws_ecs_service.url-encoder_service.network_configuration[0].subnets
}