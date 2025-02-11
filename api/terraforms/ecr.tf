resource "aws_ecr_repository" "url_encoder" {
  name         = "url-encoder"
  force_delete = true
}

resource "null_resource" "build_and_push" {
  depends_on = [aws_ecr_repository.url_encoder]

  provisioner "local-exec" {
    command = <<EOT
      # Generate a short SHA from the current git commit
      SHA=$(git rev-parse --short HEAD)
      echo "Building docker image"
      docker build -f ./api/Dockerfile -t url-encoder:latest ./api
      
      echo "Logging in to ECR"
      aws ecr get-login-password --region ${var.aws_region} | docker login --username AWS --password-stdin ${var.aws_account_id}.dkr.ecr.${var.aws_region}.amazonaws.com
      
      echo "Tagging image for ECR"
      docker tag url-encoder:latest ${var.aws_account_id}.dkr.ecr.${var.aws_region}.amazonaws.com/url-encoder:latest
      
      echo "Pushing image to ECR"
      docker push ${var.aws_account_id}.dkr.ecr.${var.aws_region}.amazonaws.com/url-encoder:latest
EOT
  }
}