# Build and Deployment Guide

## Quick Build

```bash
# Install dependencies
go mod download

# Build for current platform
make build

# Build for all platforms
make build-all

# Install locally
make install
```

## Manual Build Commands

### Linux
```bash
GOOS=linux GOARCH=amd64 go build -o bin/sniprun-linux-amd64
```

### macOS (Intel)
```bash
GOOS=darwin GOARCH=amd64 go build -o bin/sniprun-darwin-amd64
```

### macOS (Apple Silicon)
```bash
GOOS=darwin GOARCH=arm64 go build -o bin/sniprun-darwin-arm64
```

### Windows
```bash
GOOS=windows GOARCH=amd64 go build -o bin/sniprun-windows-amd64.exe
```

## Project Structure

```
sniprun/
├── main.go                  # Entry point
├── go.mod                   # Dependencies
├── Makefile                 # Build automation
├── README.md                # User documentation
├── BUILD_AND_DEPLOY.md      # This file
├── cmd/                     # CLI commands
│   ├── root.go             # CLI root
│   ├── run.go              # Execute snips
│   ├── list.go             # List snips
│   ├── add.go              # Create snips
│   ├── explain.go          # Show snip details
│   ├── remove.go           # Delete snips
│   └── update.go           # Sync community snips
├── internal/                # Internal packages
│   ├── snip/
│   │   ├── snip.go         # Snip data model
│   │   └── executor.go     # Command execution
│   ├── security/
│   │   └── validator.go    # Gemini API integration
│   └── repo/
│       └── sync.go         # GitHub sync
└── install.ps1             # Windows installer

```

## Dependencies

```bash
# Required
go >= 1.21

# Go packages (auto-installed)
github.com/spf13/cobra      # CLI framework
gopkg.in/yaml.v3            # YAML parsing
```

## Environment Variables

```bash
# Optional: Enable security validation
export GEMINI_API_KEY="your-gemini-api-key"

# Optional: Custom config directory
export SNIPRUN_CONFIG_DIR="$HOME/.sniprun"
```

## Distribution

### 1. GitHub Releases

Create a release workflow (`.github/workflows/release.yml`):

```yaml
name: Release

on:
  push:
    tags:
      - 'v*'

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      
      - name: Build all platforms
        run: make build-all
      
      - name: Create Release
        uses: softprops/action-gh-release@v1
        with:
          files: bin/*
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

### 2. Homebrew (macOS/Linux)

Create a Homebrew tap:

```ruby
# Formula/sniprun.rb
class Sniprun < Formula
  desc "Run complex commands with short, memorable snips"
  homepage "https://github.com/yourusername/sniprun"
  url "https://github.com/yourusername/sniprun/archive/v0.1.0.tar.gz"
  sha256 "..."
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(ldflags: "-s -w")
  end

  test do
    system "#{bin}/sniprun", "--version"
  end
end
```

Install with:
```bash
brew tap yourusername/sniprun
brew install sniprun
```

### 3. Winget (Windows)

Create a manifest for Windows Package Manager:

```yaml
# manifests/s/sniprun/sniprun/0.1.0/sniprun.sniprun.yaml
PackageIdentifier: sniprun.sniprun
PackageVersion: 0.1.0
PackageLocale: en-US
Publisher: Your Name
PackageName: sniprun
License: MIT
ShortDescription: Run complex commands with short, memorable snips
Installers:
  - Architecture: x64
    InstallerType: portable
    InstallerUrl: https://github.com/yourusername/sniprun/releases/download/v0.1.0/sniprun-windows-amd64.exe
    InstallerSha256: ...
```

Install with:
```powershell
winget install sniprun.sniprun
```

### 4. AUR (Arch Linux)

Create a PKGBUILD:

```bash
# PKGBUILD
pkgname=sniprun
pkgver=0.1.0
pkgrel=1
pkgdesc="Run complex commands with short, memorable snips"
arch=('x86_64')
url="https://github.com/yourusername/sniprun"
license=('MIT')
depends=()
makedepends=('go')
source=("$pkgname-$pkgver.tar.gz::$url/archive/v$pkgver.tar.gz")
sha256sums=('...')

build() {
  cd "$pkgname-$pkgver"
  go build -o sniprun
}

package() {
  cd "$pkgname-$pkgver"
  install -Dm755 sniprun "$pkgdir/usr/bin/sniprun"
}
```

### 5. Docker

```dockerfile
# Dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o sniprun

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/sniprun /usr/local/bin/
ENTRYPOINT ["sniprun"]
```

Build and run:
```bash
docker build -t sniprun .
docker run -it sniprun list
```

## Testing

### Unit Tests

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific package
go test ./internal/snip/
```

### Integration Tests

Create `test/integration_test.go`:

```go
package test

import (
    "os"
    "testing"
    "sniprun/internal/snip"
)

func TestSnipExecution(t *testing.T) {
    // Test snip loading
    s := &snip.Snip{
        Name:    "test",
        Command: "echo hello",
        Args:    []string{},
    }
    
    err := s.Execute([]string{}, false)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
}
```

### Manual Testing Checklist

- [ ] `sniprun list` shows snips
- [ ] `sniprun add test-snip` creates snip
- [ ] `sniprun test-snip` executes successfully
- [ ] `sniprun explain test-snip` shows details
- [ ] `sniprun remove test-snip` deletes snip
- [ ] `sniprun update` syncs from GitHub
- [ ] Security validation works (with GEMINI_API_KEY)
- [ ] Argument interpolation works correctly
- [ ] Cross-platform execution (Windows PowerShell, Unix sh)

## Community Repository Setup

### 1. Create Snips Repository

```bash
# Create new repo: github.com/sniprun/snips
mkdir snips
cd snips
git init

# Organize by category
mkdir -p docker git nodejs python system windows powershell

# Add example snips
cat > docker/docker-clean.yaml << EOF
name: docker-clean
description: Remove all Docker containers, images, volumes
command: docker system prune -a --volumes -f
args: []
category: docker
trust: community
EOF

# Add README
cat > README.md << EOF
# sniprun Community Snips

Community-contributed command snippets for sniprun.

## Contributing

1. Fork this repository
2. Add your snip in the appropriate category folder
3. Test locally: \`sniprun your-snip\`
4. Submit a pull request

## Snip Guidelines

- Keep commands focused and reusable
- Add clear descriptions
- Test on multiple platforms when possible
- Use meaningful argument names
EOF

git add .
git commit -m "Initial snips"
git push origin main
```

### 2. CI/CD for Snip Validation

`.github/workflows/validate.yml`:

```yaml
name: Validate Snips

on: [pull_request]

jobs:
  validate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Validate YAML
        run: |
          for file in **/*.yaml; do
            echo "Validating $file"
            python3 -c "import yaml; yaml.safe_load(open('$file'))"
          done
      
      - name: Check required fields
        run: |
          for file in **/*.yaml; do
            echo "Checking $file"
            grep -q "name:" "$file" || exit 1
            grep -q "description:" "$file" || exit 1
            grep -q "command:" "$file" || exit 1
          done
```

## Release Checklist

- [ ] Update version in `go.mod`
- [ ] Update CHANGELOG.md
- [ ] Run full test suite: `make test`
- [ ] Build for all platforms: `make build-all`
- [ ] Test on Windows, macOS, Linux
- [ ] Create git tag: `git tag v0.1.0`
- [ ] Push tag: `git push origin v0.1.0`
- [ ] Create GitHub release with binaries
- [ ] Update Homebrew formula
- [ ] Update Winget manifest
- [ ] Announce on social media

## Troubleshooting

### Build Errors

**"package not found"**
```bash
go mod download
go mod tidy
```

**"permission denied"**
```bash
chmod +x bin/sniprun
```

### Runtime Errors

**"snip not found"**
```bash
# Update community snips
sniprun update

# Check config directory
ls ~/.sniprun/snips/
```

**"command not found: sniprun"**
```bash
# Check PATH
echo $PATH

# Reinstall
make install
```

## Performance Optimization

### Binary Size Reduction

```bash
# Strip debug symbols
go build -ldflags="-s -w" -o sniprun

# Use UPX compression (optional)
upx --best --lzma sniprun
```

### Startup Time

- Lazy load snips (only when needed)
- Cache parsed YAML in memory
- Use goroutines for parallel operations

## Security Best Practices

1. **Never execute untrusted snips without review**
2. **Enable Gemini API validation in production**
3. **Review community snips before adding**
4. **Use `--skip-check` only for trusted commands**
5. **Keep local snips separate from community**

## Contributing to Core

1. Fork the repository
2. Create feature branch: `git checkout -b feature/amazing-feature`
3. Make changes and test thoroughly
4. Commit: `git commit -m 'Add amazing feature'`
5. Push: `git push origin feature/amazing-feature`
6. Open Pull Request

## Support

- GitHub Issues: Report bugs and feature requests
- Discussions: Ask questions and share ideas
- Wiki: Extended documentation and guides