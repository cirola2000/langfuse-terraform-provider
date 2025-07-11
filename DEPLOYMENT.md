# Langfuse Terraform Provider - Deployment Guide

This guide covers how to deploy and use the Langfuse Terraform Provider across different platforms.

## Quick Start

### 1. Download Pre-built Binaries

Choose the appropriate binary for your platform:

| Platform | Architecture | Download | Install Command |
|----------|-------------|----------|-----------------|
| macOS | Intel (x86_64) | `dist/terraform-provider-langfuse_1.0.0_darwin_amd64.zip` | `make install` |
| macOS | Apple Silicon (ARM64) | `dist/terraform-provider-langfuse_1.0.0_darwin_arm64.zip` | `make install` |
| Linux | x86_64 | `dist/terraform-provider-langfuse_1.0.0_linux_amd64.zip` | `./install-linux.sh` |
| Linux | ARM64 | `dist/terraform-provider-langfuse_1.0.0_linux_arm64.zip` | `./install-linux.sh` |
| Windows | x86_64 | `dist/terraform-provider-langfuse_1.0.0_windows_amd64.zip` | Manual install |

### 2. Installation on Linux

```bash
# Extract the downloaded zip file
unzip terraform-provider-langfuse_1.0.0_linux_amd64.zip

# Run the installation script
./install-linux.sh
```

### 3. Cross-Platform Development

If you're developing on macOS/Windows but deploying on Linux:

```bash
# Build for Linux from any platform
make build-all

# Or build specific platform
GOOS=linux GOARCH=amd64 go build -o terraform-provider-langfuse-linux-amd64

# Transfer to Linux server
scp terraform-provider-langfuse-linux-amd64 user@server:/home/user/

# On Linux server, install manually:
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/cirobaronneto/langfuse/1.0.0/linux_amd64
cp terraform-provider-langfuse-linux-amd64 ~/.terraform.d/plugins/registry.terraform.io/cirobaronneto/langfuse/1.0.0/linux_amd64/terraform-provider-langfuse
chmod +x ~/.terraform.d/plugins/registry.terraform.io/cirobaronneto/langfuse/1.0.0/linux_amd64/terraform-provider-langfuse
```

## Verification

After installation, verify the provider works:

```bash
cd examples/
terraform init
terraform validate
```

You should see:
```
Terraform has been successfully initialized!
Success! The configuration is valid.
```

## CI/CD Integration

### GitHub Actions Example

```yaml
name: Deploy with Langfuse Provider

on:
  push:
    branches: [main]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Download Langfuse Provider
        run: |
          wget https://github.com/your-repo/releases/download/v1.0.0/terraform-provider-langfuse_1.0.0_linux_amd64.zip
          unzip terraform-provider-langfuse_1.0.0_linux_amd64.zip
          
      - name: Install Provider
        run: |
          mkdir -p ~/.terraform.d/plugins/registry.terraform.io/cirobaronneto/langfuse/1.0.0/linux_amd64
          cp terraform-provider-langfuse-linux-amd64 ~/.terraform.d/plugins/registry.terraform.io/cirobaronneto/langfuse/1.0.0/linux_amd64/terraform-provider-langfuse
          chmod +x ~/.terraform.d/plugins/registry.terraform.io/cirobaronneto/langfuse/1.0.0/linux_amd64/terraform-provider-langfuse
          
      - name: Setup Terraform
        uses: hashicorp/setup-terraform@v2
        
      - name: Terraform Deploy
        run: |
          terraform init
          terraform plan
          terraform apply -auto-approve
        env:
          LANGFUSE_API_HOST: https://cloud.langfuse.com
          LANGFUSE_SECRET_KEY: ${{ secrets.LANGFUSE_SECRET_KEY }}
          LANGFUSE_PUBLIC_KEY: ${{ secrets.LANGFUSE_PUBLIC_KEY }}
```

### Docker Example

```dockerfile
FROM hashicorp/terraform:latest

# Install Langfuse provider
COPY terraform-provider-langfuse-linux-amd64 /usr/local/bin/terraform-provider-langfuse
RUN chmod +x /usr/local/bin/terraform-provider-langfuse && \
    mkdir -p /root/.terraform.d/plugins/registry.terraform.io/cirobaronneto/langfuse/1.0.0/linux_amd64 && \
    cp /usr/local/bin/terraform-provider-langfuse /root/.terraform.d/plugins/registry.terraform.io/cirobaronneto/langfuse/1.0.0/linux_amd64/

WORKDIR /terraform
COPY . .

RUN terraform init
CMD ["terraform", "apply", "-auto-approve"]
```

## Troubleshooting

### Provider Not Found

If you get "provider not found" errors:

1. Check the provider is in the correct directory:
   ```bash
   ls -la ~/.terraform.d/plugins/registry.terraform.io/cirobaronneto/langfuse/1.0.0/
   ```

2. Verify the binary is executable:
   ```bash
   chmod +x ~/.terraform.d/plugins/registry.terraform.io/cirobaronneto/langfuse/1.0.0/*/terraform-provider-langfuse
   ```

3. Check platform compatibility:
   ```bash
   file ~/.terraform.d/plugins/registry.terraform.io/cirobaronneto/langfuse/1.0.0/*/terraform-provider-langfuse
   ```

### Authentication Issues

Ensure your Langfuse credentials are properly set:

```bash
export LANGFUSE_API_HOST="https://cloud.langfuse.com"
export LANGFUSE_SECRET_KEY="sk-lf-your-secret-key"
export LANGFUSE_PUBLIC_KEY="pk-lf-your-public-key"
```

Or configure in your Terraform files:

```hcl
provider "langfuse" {
  api_host   = "https://cloud.langfuse.com"
  secret_key = var.langfuse_secret_key
  public_key = var.langfuse_public_key
}
```

## Building from Source

If you need to build from source on your target platform:

```bash
# Install Go (if not already installed)
# Ubuntu/Debian: sudo apt update && sudo apt install golang-go
# RHEL/CentOS: sudo yum install golang
# macOS: brew install go

# Clone and build
git clone <repository-url>
cd langfuse-terraform-provider
make install
```

## Support

For issues and questions:
- Check the provider logs: `TF_LOG=DEBUG terraform plan`
- Verify Langfuse API connectivity: `curl -u pk:sk https://cloud.langfuse.com/api/public/organizations/projects`
- Review Terraform provider documentation 