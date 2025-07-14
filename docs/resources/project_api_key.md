# langfuse_project_api_key Resource

Manages a Langfuse project API key. API keys are used to authenticate applications and services with the Langfuse API for a specific project.

## Example Usage

### Basic Usage

```hcl
resource "langfuse_project" "example" {
  name = "my-project"
}

resource "langfuse_project_api_key" "example" {
  project_id = langfuse_project.example.id
}
```

### API Key with Note

```hcl
resource "langfuse_project" "example" {
  name = "production-app"
}

resource "langfuse_project_api_key" "production" {
  project_id = langfuse_project.example.id
  note       = "Production environment API key"
}
```

### Multiple API Keys for Different Environments

```hcl
resource "langfuse_project" "app" {
  name = "my-application"
  
  metadata = {
    environment = "production"
  }
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

### Using the API Keys in Application Configuration

```hcl
resource "langfuse_project" "example" {
  name = "my-project"
}

resource "langfuse_project_api_key" "app_key" {
  project_id = langfuse_project.example.id
  note       = "Application API key"
}

# Example: Use the API key in other resources or outputs
output "langfuse_config" {
  value = {
    public_key  = langfuse_project_api_key.app_key.public_key
    secret_key  = langfuse_project_api_key.app_key.secret_key
    project_id  = langfuse_project.example.id
  }
  sensitive = true
}
```

## Schema

### Required

- `project_id` (String) The ID of the project that the API key belongs to. This field requires replacement if changed.

### Optional

- `note` (String) Optional note or description for the API key. Useful for identifying the purpose or environment of the key. This field requires replacement if changed.

### Read-Only

- `id` (String) The unique identifier of the API key.
- `public_key` (String) The public key portion of the API key. This is used as the username in basic authentication.
- `secret_key` (String, Sensitive) The secret key portion of the API key. This is only available immediately after creation and is used as the password in basic authentication.
- `display_secret_key` (String) A partial display version of the secret key for identification purposes.
- `created_at` (String) The timestamp when the API key was created (RFC3339 format).
- `expires_at` (String) The timestamp when the API key expires (RFC3339 format). May be null if the key doesn't expire.
- `last_used_at` (String) The timestamp when the API key was last used (RFC3339 format). May be null if the key has never been used.

## Important Notes

### API Key Immutability

API keys are **immutable** after creation. Any changes to the `project_id` or `note` will force Terraform to destroy and recreate the resource. This is by design to maintain security best practices.

### Secret Key Availability

The `secret_key` is only returned by the Langfuse API during creation. After that, only the `display_secret_key` (a partial key for identification) is available. Make sure to:

1. **Store the secret key securely** in your Terraform state
2. **Use remote state** with encryption for production environments
3. **Rotate keys regularly** for security

### Authentication Usage

Use the API key for authentication with Langfuse APIs:

```bash
# Example using curl
curl -X GET "https://cloud.langfuse.com/api/public/projects" \
  -u "$PUBLIC_KEY:$SECRET_KEY"
```

## Import

Import is not currently supported for API key resources due to the sensitive nature of the secret key. If you need to manage existing API keys with Terraform, you'll need to recreate them.

## Security Considerations

1. **State File Security**: API keys are stored in the Terraform state file. Ensure your state is encrypted and stored securely.

2. **Access Control**: Limit access to Terraform state files and the systems that run Terraform.

3. **Key Rotation**: Regularly rotate API keys and update your applications accordingly.

4. **Environment Separation**: Use different API keys for different environments (development, staging, production).

5. **Monitoring**: Monitor API key usage through the `last_used_at` attribute and Langfuse's audit logs. 