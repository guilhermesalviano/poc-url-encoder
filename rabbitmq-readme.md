# RabbitMQ Integration Guide

## Prerequisites
- Docker
- Docker Compose
- Go 1.19+

## Running RabbitMQ

### Start Services
```bash
docker-compose up -d
```

### Check RabbitMQ Status
```bash
# Check running containers
docker-compose ps

# View RabbitMQ logs
docker-compose logs rabbitmq
```

## Testing RabbitMQ

### Management UI
- URL: http://localhost:15672
- Default Credentials:
  - Username: guest
  - Password: guest

### Verification Commands
```bash
# Check RabbitMQ is running
docker exec -it <container_name> rabbitmqctl status

# List queues
docker exec -it <container_name> rabbitmqctl list_queues
```

## Go Dependencies
```bash
# Install AMQP library
go get github.com/streadway/amqp
```

## Connection Configuration
```go
uri := "amqp://guest:guest@localhost:5672/"
```

## Common Issues
- Ensure port 5672 and 15672 are not in use
- Check Docker network configuration
- Verify RabbitMQ credentials

## Troubleshooting
1. Restart services
```bash
docker-compose down
docker-compose up -d
```

2. Check container logs
```bash
docker-compose logs rabbitmq
```
