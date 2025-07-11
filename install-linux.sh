#!/bin/bash
set -e

# Langfuse Terraform Provider Linux Installation Script

PROVIDER_VERSION="1.0.0"
PROVIDER_NAME="terraform-provider-langfuse"
REGISTRY_PATH="registry.terraform.io/cirola2000/langfuse"

# Detect architecture
ARCH=$(uname -m)
case $ARCH in
    x86_64)
        ARCH="amd64"
        ;;
    aarch64)
        ARCH="arm64"
        ;;
    arm64)
        ARCH="arm64"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

PLATFORM="linux_${ARCH}"
PLUGIN_DIR="$HOME/.terraform.d/plugins/${REGISTRY_PATH}/${PROVIDER_VERSION}/${PLATFORM}"
BINARY_NAME="${PROVIDER_NAME}-linux-${ARCH}"

echo "Installing Langfuse Terraform Provider for ${PLATFORM}..."

# Create plugin directory
mkdir -p "$PLUGIN_DIR"

# Check if binary exists
if [ ! -f "$BINARY_NAME" ]; then
    echo "Error: Binary $BINARY_NAME not found in current directory"
    echo "Available binaries:"
    ls -la terraform-provider-langfuse-* 2>/dev/null || echo "No binaries found"
    echo ""
    echo "To build the binary, run:"
    echo "  GOOS=linux GOARCH=${ARCH} go build -o ${BINARY_NAME}"
    exit 1
fi

# Copy binary
cp "$BINARY_NAME" "$PLUGIN_DIR/terraform-provider-langfuse"
chmod +x "$PLUGIN_DIR/terraform-provider-langfuse"

echo "✅ Successfully installed Langfuse Terraform Provider to:"
echo "   $PLUGIN_DIR/terraform-provider-langfuse"
echo ""
echo "You can now use the provider in your Terraform configurations:"
echo ""
echo "terraform {"
echo "  required_providers {"
echo "    langfuse = {"
echo "      source = \"${REGISTRY_PATH}\""
echo "    }"
echo "  }"
echo "}"
echo ""
echo "provider \"langfuse\" {"
echo "  api_host   = \"https://cloud.langfuse.com\""
echo "  secret_key = var.langfuse_secret_key"
echo "  public_key = var.langfuse_public_key"
echo "}" 