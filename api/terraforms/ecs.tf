resource "aws_ecs_cluster" "url-encoder_cluster" {
  name = "url-encoder-cluster"

  depends_on = [null_resource.build_and_push]
}

resource "aws_ecs_task_definition" "url-encoder_task" {
  family                   = "url-encoder"
  requires_compatibilities = ["FARGATE"]
  network_mode            = "awsvpc"
  cpu                     = "256"
  memory                  = "512"
  execution_role_arn      = aws_iam_role.ecs_execution_role.arn

  container_definitions = jsonencode([{
    name      = "url-encoder"
    image     = aws_ecr_repository.url_encoder.repository_url
    cpu       = 256
    memory    = 512
    essential = true

    portMappings = [{
      containerPort = 3333
      hostPort      = 3333
      protocol      = "tcp"
    }]

    logConfiguration = {
      logDriver = "awslogs"
      options = {
        "awslogs-group"         = aws_cloudwatch_log_group.ecs_log_group.name
        "awslogs-region"        = var.aws_region
        "awslogs-stream-prefix" = "ecs"
      }
    }
  }])
}

# Create ECS Service
resource "aws_ecs_service" "url-encoder_service" {
  name            = "url-encoder-service"
  cluster         = aws_ecs_cluster.url-encoder_cluster.id
  task_definition = aws_ecs_task_definition.url-encoder_task.id
  desired_count   = 1
  launch_type     = "FARGATE"
  enable_ecs_managed_tags = true

  network_configuration {
    subnets          = [aws_subnet.my_public_subnet.id]
    security_groups  = [aws_security_group.my_security_group.id]
    assign_public_ip = true
  }
}
