terraform {
  required_providers {
    langfuse = {
      source = "registry.terraform.io/cirola2000/langfuse"
    }
  }
}

# Configure the Langfuse Provider
provider "langfuse" {
  api_host   = "https://cloud.langfuse.com" # or use LANGFUSE_API_HOST env var
  secret_key = var.langfuse_secret_key      # or use LANGFUSE_SECRET_KEY env var
  public_key = var.langfuse_public_key      # or use LANGFUSE_PUBLIC_KEY env var
}

# Variables for sensitive data
variable "langfuse_secret_key" {
  description = "Langfuse Secret Key"
  type        = string
  sensitive   = true
}

variable "langfuse_public_key" {
  description = "Langfuse Public Key"
  type        = string
}

# Create a Langfuse project
resource "langfuse_project" "example" {
  name = "my-terraform-project"

  metadata = {
    environment = "production"
    team        = "data-team"
    cost_center = "engineering"
  }

  retention_days = 30
}

# Create API keys for the project
resource "langfuse_project_api_key" "production" {
  project_id = langfuse_project.example.id
  note       = "Production API key"
}

resource "langfuse_project_api_key" "monitoring" {
  project_id = langfuse_project.example.id
  note       = "Monitoring and analytics"
}

# Output the project ID
output "project_id" {
  value = langfuse_project.example.id
}

output "project_created_at" {
  value = langfuse_project.example.created_at
}

output "project_updated_at" {
  value = langfuse_project.example.updated_at
}

# Output API key information (sensitive)
output "production_api_key" {
  value = {
    public_key = langfuse_project_api_key.production.public_key
    secret_key = langfuse_project_api_key.production.secret_key
  }
  sensitive = true
}

output "monitoring_api_key_public" {
  value       = langfuse_project_api_key.monitoring.public_key
  description = "Public key for monitoring API key (safe to share)"
}
