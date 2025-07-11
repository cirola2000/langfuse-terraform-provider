# Langfuse Provider Examples

This directory contains example Terraform configurations using the Langfuse provider.

## Setup

1. Copy `terraform.tfvars.example` to `terraform.tfvars`:
   ```bash
   cp terraform.tfvars.example terraform.tfvars
   ```

2. Edit `terraform.tfvars` and add your actual Langfuse API credentials:
   ```hcl
   langfuse_secret_key = "sk-lf-your-actual-secret-key"
   langfuse_public_key = "pk-lf-your-actual-public-key"
   ```

3. Initialize and apply:
   ```bash
   terraform init
   terraform plan
   terraform apply
   ```

## Getting Langfuse API Keys

1. Go to [Langfuse Cloud](https://cloud.langfuse.com)
2. Sign up or log in to your account
3. Navigate to Settings → API Keys
4. Create an organization-scoped API key
5. Copy the public and secret keys

## Environment Variables

Alternatively, you can set environment variables instead of using `terraform.tfvars`:

```bash
export LANGFUSE_API_HOST="https://cloud.langfuse.com"
export LANGFUSE_SECRET_KEY="sk-lf-your-secret-key"
export LANGFUSE_PUBLIC_KEY="pk-lf-your-public-key"
```

## What This Example Does

The main.tf file demonstrates:

- Configuring the Langfuse provider
- Creating a project with metadata and retention settings
- Outputting project information

After applying, you'll see outputs showing the project ID and timestamps. 