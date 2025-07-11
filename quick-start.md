# Quick Start: Auto-Download Provider

This guide shows you how to use the Langfuse Terraform Provider with **automatic downloads** - no manual installation required!

## 🚀 One-Command Setup

### Option 1: Universal Setup Script (Recommended)

```bash
# Download and run the setup script
curl -fsSL https://raw.githubusercontent.com/cirola2000/langfuse-terraform-provider/main/terraform-mirror-setup.sh | bash
```

This script will:
- Detect your platform automatically
- Download the correct provider binary
- Set up Terraform to find it automatically
- No manual installation steps needed!

### Option 2: Manual Setup (if you prefer)

<details>
<summary>Click to expand manual steps</summary>

```bash
# 1. Detect your platform
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)
case $ARCH in x86_64) ARCH="amd64";; aarch64|arm64) ARCH="arm64";; esac

# 2. Create provider directory
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/cirola2000/langfuse/1.0.0/${OS}_${ARCH}

# 3. Download and install
cd /tmp
curl -L -o provider.zip "https://github.com/cirola2000/langfuse-terraform-provider/raw/main/dist/terraform-provider-langfuse_1.0.0_${OS}_${ARCH}.zip"
unzip provider.zip
chmod +x terraform-provider-langfuse-*
cp terraform-provider-langfuse-* ~/.terraform.d/plugins/registry.terraform.io/cirola2000/langfuse/1.0.0/${OS}_${ARCH}/terraform-provider-langfuse_v1.0.0
```

</details>

## 📝 Create Your Terraform Configuration

Create a `main.tf` file:

```hcl
terraform {
  required_providers {
    langfuse = {
      source  = "cirola2000/langfuse"
      version = "1.0.0"
    }
  }
}

provider "langfuse" {
  api_host   = "https://cloud.langfuse.com"
  secret_key = var.langfuse_secret_key
  public_key = var.langfuse_public_key
}

variable "langfuse_secret_key" {
  description = "Langfuse Secret Key"
  type        = string
  sensitive   = true
}

variable "langfuse_public_key" {
  description = "Langfuse Public Key"
  type        = string
}

resource "langfuse_project" "example" {
  name = "terraform-managed-project"
  
  metadata = {
    environment = "production"
    team        = "engineering"
    managed_by  = "terraform"
  }
  
  retention_days = 30
}

output "project_id" {
  value = langfuse_project.example.id
}
```

Create a `terraform.tfvars` file:

```hcl
langfuse_secret_key = "sk-lf-your-secret-key-here"
langfuse_public_key = "pk-lf-your-public-key-here"
```

## 🎯 Deploy

```bash
# Initialize - Terraform will automatically find the provider!
terraform init

# Plan your changes
terraform plan

# Apply your infrastructure
terraform apply
```

## ✨ What Just Happened?

1. **Automatic Discovery**: Terraform automatically found your provider in the local plugin directory
2. **No Downloads**: No need to manually download or install anything during `terraform init`
3. **Version Locking**: Terraform uses exactly version 1.0.0 as specified
4. **Platform Detection**: The setup script automatically chose the right binary for your system

## 🔧 Environment Variables (Alternative)

Instead of `terraform.tfvars`, you can use environment variables:

```bash
export LANGFUSE_API_HOST="https://cloud.langfuse.com"
export LANGFUSE_SECRET_KEY="sk-lf-your-secret-key"
export LANGFUSE_PUBLIC_KEY="pk-lf-your-public-key"

terraform apply
```

## 🐳 Docker Usage

For containerized deployments:

```dockerfile
FROM hashicorp/terraform:latest

# Install the provider
RUN curl -fsSL https://raw.githubusercontent.com/cirola2000/langfuse-terraform-provider/main/terraform-mirror-setup.sh | bash

WORKDIR /terraform
COPY . .
RUN terraform init
```

## 🎉 Success!

Your Langfuse Terraform Provider is now set up for **automatic use**! 

- ✅ No manual binary downloads needed
- ✅ Terraform automatically finds the provider
- ✅ Version-locked and platform-specific
- ✅ Works the same way as official providers
- ✅ Perfect for CI/CD pipelines

Happy Infrastructure as Code! 🚀 