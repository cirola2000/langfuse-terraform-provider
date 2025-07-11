default: install

generate:
	go generate ./...

build:
	go build -o terraform-provider-langfuse

# Build for multiple platforms
build-all:
	GOOS=darwin GOARCH=arm64 go build -o terraform-provider-langfuse-darwin-arm64
	GOOS=darwin GOARCH=amd64 go build -o terraform-provider-langfuse-darwin-amd64
	GOOS=linux GOARCH=amd64 go build -o terraform-provider-langfuse-linux-amd64
	GOOS=linux GOARCH=arm64 go build -o terraform-provider-langfuse-linux-arm64
	GOOS=windows GOARCH=amd64 go build -o terraform-provider-langfuse-windows-amd64.exe

# Install for current platform (detected automatically)
install: build
	@OS=$$(uname -s | tr '[:upper:]' '[:lower:]'); \
	ARCH=$$(uname -m); \
	case $$ARCH in \
		x86_64) ARCH=amd64 ;; \
		aarch64) ARCH=arm64 ;; \
		arm64) ARCH=arm64 ;; \
	esac; \
	mkdir -p ~/.terraform.d/plugins/registry.terraform.io/cirola2000/langfuse/1.0.0/$${OS}_$${ARCH}; \
	cp terraform-provider-langfuse ~/.terraform.d/plugins/registry.terraform.io/cirola2000/langfuse/1.0.0/$${OS}_$${ARCH}/

# Install all platforms (useful for development)
install-all: build-all
	# Darwin ARM64
	mkdir -p ~/.terraform.d/plugins/registry.terraform.io/cirola2000/langfuse/1.0.0/darwin_arm64
	cp terraform-provider-langfuse-darwin-arm64 ~/.terraform.d/plugins/registry.terraform.io/cirola2000/langfuse/1.0.0/darwin_arm64/terraform-provider-langfuse
	# Darwin AMD64
	mkdir -p ~/.terraform.d/plugins/registry.terraform.io/cirola2000/langfuse/1.0.0/darwin_amd64
	cp terraform-provider-langfuse-darwin-amd64 ~/.terraform.d/plugins/registry.terraform.io/cirola2000/langfuse/1.0.0/darwin_amd64/terraform-provider-langfuse
	# Linux AMD64
	mkdir -p ~/.terraform.d/plugins/registry.terraform.io/cirola2000/langfuse/1.0.0/linux_amd64
	cp terraform-provider-langfuse-linux-amd64 ~/.terraform.d/plugins/registry.terraform.io/cirola2000/langfuse/1.0.0/linux_amd64/terraform-provider-langfuse
	# Linux ARM64
	mkdir -p ~/.terraform.d/plugins/registry.terraform.io/cirola2000/langfuse/1.0.0/linux_arm64
	cp terraform-provider-langfuse-linux-arm64 ~/.terraform.d/plugins/registry.terraform.io/cirola2000/langfuse/1.0.0/linux_arm64/terraform-provider-langfuse
	# Windows AMD64
	mkdir -p ~/.terraform.d/plugins/registry.terraform.io/cirola2000/langfuse/1.0.0/windows_amd64
	cp terraform-provider-langfuse-windows-amd64.exe ~/.terraform.d/plugins/registry.terraform.io/cirola2000/langfuse/1.0.0/windows_amd64/terraform-provider-langfuse.exe

# Create distribution packages
dist: build-all
	mkdir -p dist
	# Darwin ARM64
	zip dist/terraform-provider-langfuse_1.0.0_darwin_arm64.zip terraform-provider-langfuse-darwin-arm64
	# Darwin AMD64  
	zip dist/terraform-provider-langfuse_1.0.0_darwin_amd64.zip terraform-provider-langfuse-darwin-amd64
	# Linux AMD64
	zip dist/terraform-provider-langfuse_1.0.0_linux_amd64.zip terraform-provider-langfuse-linux-amd64
	# Linux ARM64
	zip dist/terraform-provider-langfuse_1.0.0_linux_arm64.zip terraform-provider-langfuse-linux-arm64
	# Windows AMD64
	zip dist/terraform-provider-langfuse_1.0.0_windows_amd64.zip terraform-provider-langfuse-windows-amd64.exe

test:
	go test -count=1 -parallel=4 ./...

testacc:
	TF_ACC=1 go test -count=1 -parallel=4 -timeout 10m -v ./...

fmt:
	go fmt ./...

clean:
	rm -f terraform-provider-langfuse*
	rm -rf dist/

.PHONY: build build-all install install-all dist test testacc fmt clean generate 