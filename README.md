# Terraform Provider for Langfuse

Manage [Langfuse](https://langfuse.com/) projects and API keys with Terraform.

## Quick Start

```hcl
terraform {
  required_providers {
    langfuse = {
      source  = "cirola2000/langfuse"
      version = "~> 1.1"
    }
  }
}

provider "langfuse" {
  api_host   = "https://cloud.langfuse.com"
  secret_key = var.langfuse_secret_key
  public_key = var.langfuse_public_key
}

# Create a project
resource "langfuse_project" "example" {
  name = "my-project"
  
  metadata = {
    environment = "production"
    team        = "data-team"
  }
  
  retention_days = 30
}

# Create API keys for the project
resource "langfuse_project_api_key" "production" {
  project_id = langfuse_project.example.id
  note       = "Production API key"
}
```

## Authentication

Set your Langfuse credentials as environment variables:

```bash
export LANGFUSE_API_HOST="https://cloud.langfuse.com"
export LANGFUSE_SECRET_KEY="sk-lf-..."
export LANGFUSE_PUBLIC_KEY="pk-lf-..."
```

Or use Terraform variables:

```hcl
variable "langfuse_secret_key" {
  description = "Langfuse Secret Key"
  type        = string
  sensitive   = true
}

variable "langfuse_public_key" {
  description = "Langfuse Public Key"
  type        = string
}
```

## Features

- **Projects**: Create and manage Langfuse projects with metadata and retention settings
- **API Keys**: Generate and manage project-specific API keys for authentication
- **Environment Variables**: Support for configuration via environment variables
- **State Management**: Secure handling of sensitive API keys

## Resources

| Resource | Description |
|----------|-------------|
| `langfuse_project` | Manage Langfuse projects |
| `langfuse_project_api_key` | Manage project API keys |

## Examples

### Basic Project

```hcl
resource "langfuse_project" "basic" {
  name = "my-langfuse-project"
}
```

### Project with Metadata

```hcl
resource "langfuse_project" "advanced" {
  name = "production-app"
  
  metadata = {
    environment   = "production"
    team          = "ml-team"
    cost_center   = "engineering"
    application   = "recommendation-engine"
  }
  
  retention_days = 90
}
```

### Multi-Environment Setup

```hcl
resource "langfuse_project" "app" {
  name = "my-application"
}

resource "langfuse_project_api_key" "production" {
  project_id = langfuse_project.app.id
  note       = "Production environment"
}

resource "langfuse_project_api_key" "staging" {
  project_id = langfuse_project.app.id
  note       = "Staging environment"
}

resource "langfuse_project_api_key" "development" {
  project_id = langfuse_project.app.id
  note       = "Development environment"
}
```

## Documentation

- [Provider Documentation](https://registry.terraform.io/providers/cirola2000/langfuse/latest/docs)
- [Project Resource](https://registry.terraform.io/providers/cirola2000/langfuse/latest/docs/resources/project)
- [Project API Key Resource](https://registry.terraform.io/providers/cirola2000/langfuse/latest/docs/resources/project_api_key)

## Requirements

- Terraform >= 1.0
- Go >= 1.21 (for development)

## License

MIT License - see [LICENSE](LICENSE) file for details. 