# Langfuse Provider

The Langfuse provider allows you to manage Langfuse projects using Terraform. This provider enables Infrastructure as Code for Langfuse project management.

## Example Usage

```hcl
terraform {
  required_providers {
    langfuse = {
      source = "cirola2000/langfuse"
      version = "~> 1.0"
    }
  }
}

provider "langfuse" {
  api_host   = "https://cloud.langfuse.com"
  secret_key = var.langfuse_secret_key
  public_key = var.langfuse_public_key
}

resource "langfuse_project" "example" {
  name = "my-project"
  
  metadata = {
    environment = "production"
    team        = "data-team"
  }
  
  retention_days = 30
}
```

## Authentication

The provider supports authentication through:

1. **Provider configuration** (recommended for CI/CD):
```hcl
provider "langfuse" {
  api_host   = "https://cloud.langfuse.com"
  secret_key = "sk-lf-..."
  public_key = "pk-lf-..."
}
```

2. **Environment variables**:
- `LANGFUSE_API_HOST` - Langfuse API host URL
- `LANGFUSE_SECRET_KEY` - Langfuse secret key
- `LANGFUSE_PUBLIC_KEY` - Langfuse public key

## Schema

### Required

- `secret_key` (String, Sensitive) The Langfuse secret key for authentication
- `public_key` (String) The Langfuse public key for authentication

### Optional

- `api_host` (String) The Langfuse API host URL. Defaults to `https://cloud.langfuse.com` 