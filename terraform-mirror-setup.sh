#!/bin/bash
set -e

# Terraform Provider Mirror Setup Script
# This allows Terraform to automatically download the provider from GitHub

PROVIDER_NAMESPACE="cirola2000"
PROVIDER_NAME="langfuse"
PROVIDER_VERSION="1.0.0"
GITHUB_REPO="cirola2000/langfuse-terraform-provider"

echo "🚀 Setting up Terraform provider mirror for automatic downloads..."

# Detect platform
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
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
        echo "❌ Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

PLATFORM="${OS}_${ARCH}"
echo "📦 Detected platform: $PLATFORM"

# Create mirror directory structure
MIRROR_DIR="$HOME/.terraform.d/plugins/registry.terraform.io/${PROVIDER_NAMESPACE}/${PROVIDER_NAME}/${PROVIDER_VERSION}/${PLATFORM}"
mkdir -p "$MIRROR_DIR"

# Download the appropriate binary
BINARY_NAME="terraform-provider-${PROVIDER_NAME}_${PROVIDER_VERSION}_${PLATFORM}.zip"
DOWNLOAD_URL="https://github.com/${GITHUB_REPO}/raw/main/dist/${BINARY_NAME}"

echo "⬇️  Downloading provider binary..."
echo "   From: $DOWNLOAD_URL"
echo "   To: $MIRROR_DIR"

# Download and extract
cd /tmp
curl -L -o "$BINARY_NAME" "$DOWNLOAD_URL" || {
    echo "❌ Failed to download from $DOWNLOAD_URL"
    echo "   Trying alternative binary name..."
    
    # Try with alternative naming
    BINARY_NAME="terraform-provider-${PROVIDER_NAME}_${PROVIDER_VERSION}_${OS}_${ARCH}.zip"
    DOWNLOAD_URL="https://github.com/${GITHUB_REPO}/raw/main/dist/${BINARY_NAME}"
    curl -L -o "$BINARY_NAME" "$DOWNLOAD_URL" || {
        echo "❌ Failed to download provider binary"
        echo "   Please check that the release exists at: https://github.com/${GITHUB_REPO}/releases"
        exit 1
    }
}

# Extract binary
unzip -q "$BINARY_NAME"
chmod +x "terraform-provider-${PROVIDER_NAME}-${OS}-${ARCH}" 2>/dev/null || \
chmod +x "terraform-provider-${PROVIDER_NAME}_v${PROVIDER_VERSION}" 2>/dev/null || \
chmod +x terraform-provider-* 2>/dev/null

# Find the binary and copy it
BINARY_FILE=$(find . -name "terraform-provider-*" -type f -executable | head -1)
if [ -z "$BINARY_FILE" ]; then
    echo "❌ Could not find extracted binary"
    ls -la
    exit 1
fi

cp "$BINARY_FILE" "$MIRROR_DIR/terraform-provider-${PROVIDER_NAME}_v${PROVIDER_VERSION}"

# Clean up
rm -f "$BINARY_NAME" terraform-provider-*

echo "✅ Provider installed successfully!"
echo ""
echo "🎯 Now you can use the provider in any Terraform configuration:"
echo ""
echo "terraform {"
echo "  required_providers {"
echo "    ${PROVIDER_NAME} = {"
echo "      source  = \"${PROVIDER_NAMESPACE}/${PROVIDER_NAME}\""
echo "      version = \"${PROVIDER_VERSION}\""
echo "    }"
echo "  }"
echo "}"
echo ""
echo "provider \"${PROVIDER_NAME}\" {"
echo "  api_host   = \"https://cloud.langfuse.com\""
echo "  secret_key = var.langfuse_secret_key"
echo "  public_key = var.langfuse_public_key"
echo "}"
echo ""
echo "🔧 Terraform will now automatically find and use the provider!"
echo "   Run: terraform init" 