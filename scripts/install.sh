#!/bin/bash
# sniprun Unix/Linux/macOS Installer

set -e

echo "🚀 Installing sniprun..."

# Configuration
REPO_URL="https://github.com/yourusername/sniprun"
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="sniprun"

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$ARCH" in
    x86_64) ARCH="amd64" ;;
    arm64|aarch64) ARCH="arm64" ;;
    *) echo "❌ Unsupported architecture: $ARCH"; exit 1 ;;
esac

case "$OS" in
    linux) PLATFORM="linux" ;;
    darwin) PLATFORM="darwin" ;;
    *) echo "❌ Unsupported OS: $OS"; exit 1 ;;
esac

echo "📦 Detected platform: $PLATFORM-$ARCH"

# Download binary
DOWNLOAD_URL="$REPO_URL/releases/latest/download/$BINARY_NAME-$PLATFORM-$ARCH"
echo "📥 Downloading from: $DOWNLOAD_URL"

TMP_FILE=$(mktemp)
if command -v curl >/dev/null 2>&1; then
    curl -L -o "$TMP_FILE" "$DOWNLOAD_URL" 2>/dev/null || {
        echo "❌ Download failed. Building from source..."
        BUILD_FROM_SOURCE=1
    }
elif command -v wget >/dev/null 2>&1; then
    wget -O "$TMP_FILE" "$DOWNLOAD_URL" 2>/dev/null || {
        echo "❌ Download failed. Building from source..."
        BUILD_FROM_SOURCE=1
    }
else
    echo "❌ Neither curl nor wget found. Building from source..."
    BUILD_FROM_SOURCE=1
fi

# Build from source if download failed
if [ "$BUILD_FROM_SOURCE" = "1" ]; then
    if ! command -v go >/dev/null 2>&1; then
        echo "❌ Go is not installed. Please install Go from https://golang.org/dl/"
        exit 1
    fi
    
    echo "🔨 Building from source..."
    TEMP_DIR=$(mktemp -d)
    cd "$TEMP_DIR"
    git clone "$REPO_URL" .
    go build -o "$TMP_FILE" .
    cd -
    rm -rf "$TEMP_DIR"
fi

# Make executable
chmod +x "$TMP_FILE"

# Install (requires sudo for /usr/local/bin)
echo "📦 Installing to $INSTALL_DIR..."
if [ -w "$INSTALL_DIR" ]; then
    mv "$TMP_FILE" "$INSTALL_DIR/$BINARY_NAME"
else
    sudo mv "$TMP_FILE" "$INSTALL_DIR/$BINARY_NAME"
fi

# Verify installation
if command -v $BINARY_NAME >/dev/null 2>&1; then
    echo "✅ Installation successful!"
    
    # Initial setup
    echo "🔄 Running initial setup..."
    $BINARY_NAME update || echo "⚠️  Community snips sync failed. Run 'sniprun update' manually."
    
    echo ""
    echo "🎉 sniprun is ready!"
    echo ""
    echo "Quick start:"
    echo "  sniprun list              - Show available snips"
    echo "  sniprun add my-snip       - Create a custom snip"
    echo "  sniprun docker-clean      - Execute a snip"
    echo ""
    echo "For more help: sniprun --help"
else
    echo "❌ Installation failed. Please check $INSTALL_DIR is in your PATH"
    exit 1
fi