# Terraform Provider for Langfuse

A Terraform provider for managing [Langfuse](https://langfuse.com/) projects.

## Features

- Create, read, update, and delete Langfuse projects
- Manage project metadata and retention settings
- Full Terraform lifecycle support

## Quick Start

### 1. Download Pre-built Provider

Download the appropriate binary for your platform from [GitHub Releases](https://github.com/cirola2000/langfuse-terraform-provider/releases):

| Platform | Architecture | Download |
|----------|-------------|----------|
| macOS | Intel (x86_64) | `terraform-provider-langfuse_1.0.0_darwin_amd64.zip` |
| macOS | Apple Silicon (ARM64) | `terraform-provider-langfuse_1.0.0_darwin_arm64.zip` |
| Linux | x86_64 | `terraform-provider-langfuse_1.0.0_linux_amd64.zip` |
| Linux | ARM64 | `terraform-provider-langfuse_1.0.0_linux_arm64.zip` |
| Windows | x86_64 | `terraform-provider-langfuse_1.0.0_windows_amd64.zip` |

### 2. Install the Provider

#### On Linux:
```bash
# Download and extract
wget https://github.com/cirola2000/langfuse-terraform-provider/releases/download/v1.0.0/terraform-provider-langfuse_1.0.0_linux_amd64.zip
unzip terraform-provider-langfuse_1.0.0_linux_amd64.zip

# Install using the provided script
./install-linux.sh
```

#### On macOS:
```bash
# Download and extract
wget https://github.com/cirola2000/langfuse-terraform-provider/releases/download/v1.0.0/terraform-provider-langfuse_1.0.0_darwin_arm64.zip
unzip terraform-provider-langfuse_1.0.0_darwin_arm64.zip

# Install manually
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/cirola2000/langfuse/1.0.0/darwin_arm64
cp terraform-provider-langfuse-darwin-arm64 ~/.terraform.d/plugins/registry.terraform.io/cirola2000/langfuse/1.0.0/darwin_arm64/terraform-provider-langfuse
chmod +x ~/.terraform.d/plugins/registry.terraform.io/cirola2000/langfuse/1.0.0/darwin_arm64/terraform-provider-langfuse
```

### 3. Configure Your Terraform

Create a `main.tf` file:

```hcl
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
```

### 4. Set Up Variables

Create a `terraform.tfvars` file:

```hcl
langfuse_secret_key = "sk-lf-your-secret-key-here"
langfuse_public_key = "pk-lf-your-public-key-here"
```

### 5. Deploy

```bash
terraform init
terraform plan
terraform apply
```

## Usage

### Provider Configuration

```hcl
terraform {
  required_providers {
    langfuse = {
      source = "registry.terraform.io/cirola2000/langfuse"
    }
  }
}

provider "langfuse" {
  api_host   = "https://cloud.langfuse.com"
  secret_key = var.langfuse_secret_key
  public_key = var.langfuse_public_key
}
```

### Environment Variables

You can also configure the provider using environment variables:

- `LANGFUSE_API_HOST` - Langfuse API host URL
- `LANGFUSE_SECRET_KEY` - Your Langfuse secret key
- `LANGFUSE_PUBLIC_KEY` - Your Langfuse public key

### Resource: langfuse_project

```hcl
resource "langfuse_project" "example" {
  name = "my-project"
  
  metadata = {
    environment = "production"
    team        = "data-team"
  }
  
  retention_days = 30
}
```

#### Arguments

- `name` (Required) - The name of the project
- `metadata` (Optional) - A map of metadata key-value pairs
- `retention_days` (Optional) - Number of days to retain data. Must be 0 or at least 3 days

#### Attributes

- `id` - The project ID
- `created_at` - Project creation timestamp
- `updated_at` - Project last update timestamp

## Development

### Installing Locally

#### Automatic Installation (Detects Platform)

```bash
# Build and install for your current platform
make install

# Or build and install for all platforms
make install-all
```

#### Manual Installation

```bash
# Build the provider
go build -o terraform-provider-langfuse

# Create the plugins directory (adjust OS and ARCH for your platform)
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/cirola2000/langfuse/1.0.0/{OS}_{ARCH}/

# Copy the binary
cp terraform-provider-langfuse ~/.terraform.d/plugins/registry.terraform.io/cirola2000/langfuse/1.0.0/{OS}_{ARCH}/
```

#### Cross-Platform Installation

```bash
# Build for all platforms
make build-all

# Create distribution packages
make dist

# Install specific platform (example for Linux AMD64)
GOOS=linux GOARCH=amd64 go build -o terraform-provider-langfuse-linux-amd64
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/cirola2000/langfuse/1.0.0/linux_amd64/
cp terraform-provider-langfuse-linux-amd64 ~/.terraform.d/plugins/registry.terraform.io/cirola2000/langfuse/1.0.0/linux_amd64/terraform-provider-langfuse
```

### Testing

```bash
cd examples/
terraform init
terraform plan
terraform apply
```

## Linux Deployment

For Linux users, you have several deployment options:

### Option 1: Download Pre-built Binary

1. Download the appropriate binary for your architecture:
   - **Linux AMD64**: `terraform-provider-langfuse_1.0.0_linux_amd64.zip`
   - **Linux ARM64**: `terraform-provider-langfuse_1.0.0_linux_arm64.zip`

2. Extract and install:
   ```bash
   # Download and extract (example for AMD64)
   unzip terraform-provider-langfuse_1.0.0_linux_amd64.zip
   
   # Install using the provided script
   ./install-linux.sh
   ```

### Option 2: Cross-compile from macOS/Windows

```bash
# Cross-compile for Linux
GOOS=linux GOARCH=amd64 go build -o terraform-provider-langfuse-linux-amd64

# Transfer to Linux machine and install
scp terraform-provider-langfuse-linux-amd64 user@linux-server:/tmp/
ssh user@linux-server 'mkdir -p ~/.terraform.d/plugins/registry.terraform.io/cirola2000/langfuse/1.0.0/linux_amd64 && cp /tmp/terraform-provider-langfuse-linux-amd64 ~/.terraform.d/plugins/registry.terraform.io/cirola2000/langfuse/1.0.0/linux_amd64/terraform-provider-langfuse && chmod +x ~/.terraform.d/plugins/registry.terraform.io/cirola2000/langfuse/1.0.0/linux_amd64/terraform-provider-langfuse'
```

### Option 3: Build on Linux

```bash
# Clone repository on Linux machine
git clone https://github.com/cirola2000/langfuse-terraform-provider.git
cd langfuse-terraform-provider

# Install Go if not available
# Ubuntu/Debian: sudo apt update && sudo apt install golang-go
# RHEL/CentOS: sudo yum install golang
# Arch: sudo pacman -S go

# Build and install
make install
```

### Supported Platforms

- **macOS**: darwin_amd64, darwin_arm64
- **Linux**: linux_amd64, linux_arm64  
- **Windows**: windows_amd64

## Requirements

- Terraform >= 1.0
- Go >= 1.21 (for development)
- Langfuse organization-scoped API keys

### Installing Go

If you don't have Go installed, you can install it from [https://golang.org/dl/](https://golang.org/dl/) or using a package manager:

```bash
# macOS with Homebrew
brew install go

# Ubuntu/Debian
sudo apt update && sudo apt install golang-go

# Or download from https://golang.org/dl/
```

## API Reference

This provider uses the Langfuse public API:

- `GET /api/public/organizations/projects` - List projects
- `POST /api/public/projects` - Create project
- `PUT /api/public/projects/{projectId}` - Update project
- `DELETE /api/public/projects/{projectId}` - Delete project

## Getting Langfuse API Keys

1. Go to [Langfuse Cloud](https://cloud.langfuse.com)
2. Sign up or log in to your account  
3. Navigate to Settings → API Keys
4. Create an organization-scoped API key
5. Copy the public and secret keys

## License

This project is licensed under the MIT License. 