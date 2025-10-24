# sniprun Windows PowerShell Installer

$ErrorActionPreference = "Stop"

Write-Host "Installing sniprun..." -ForegroundColor Cyan

# Configuration
$REPO_URL = "https://github.com/yourusername/sniprun"
$INSTALL_DIR = "$env:LOCALAPPDATA\sniprun"
$BINARY_NAME = "sniprun.exe"

# Create installation directory
if (!(Test-Path $INSTALL_DIR)) {
    New-Item -ItemType Directory -Path $INSTALL_DIR -Force | Out-Null
    Write-Host "Created directory: $INSTALL_DIR" -ForegroundColor Green
}

# Download or build binary
Write-Host "Downloading sniprun..." -ForegroundColor Yellow

# Try to download pre-built binary
$DOWNLOAD_URL = "$REPO_URL/releases/latest/download/sniprun-windows-amd64.exe"
$BINARY_PATH = Join-Path $INSTALL_DIR $BINARY_NAME

try {
    Invoke-WebRequest -Uri $DOWNLOAD_URL -OutFile $BINARY_PATH -ErrorAction Stop
    Write-Host "Downloaded binary successfully" -ForegroundColor Green
} catch {
    Write-Host "Pre-built binary not available. Building from source..." -ForegroundColor Yellow
    
    # Check if Go is installed
    try {
        $goVersion = go version
        Write-Host "Found: $goVersion" -ForegroundColor Green
    } catch {
        Write-Host "ERROR: Go is not installed. Please install Go from https://golang.org/dl/" -ForegroundColor Red
        exit 1
    }
    
    # Clone and build
    $tempDir = Join-Path $env:TEMP "sniprun-build"
    if (Test-Path $tempDir) {
        Remove-Item -Path $tempDir -Recurse -Force
    }
    
    git clone $REPO_URL $tempDir
    Push-Location $tempDir
    go build -o $BINARY_PATH
    Pop-Location
    
    Remove-Item -Path $tempDir -Recurse -Force
    Write-Host "Built from source successfully" -ForegroundColor Green
}

# Add to PATH if not already there
$userPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($userPath -notlike "*$INSTALL_DIR*") {
    Write-Host "Adding to PATH..." -ForegroundColor Yellow
    [Environment]::SetEnvironmentVariable(
        "Path",
        "$userPath;$INSTALL_DIR",
        "User"
    )
    $env:Path = "$env:Path;$INSTALL_DIR"
    Write-Host "Added to PATH. Restart your terminal for changes to take effect." -ForegroundColor Green
} else {
    Write-Host "Already in PATH" -ForegroundColor Green
}

# Test installation
Write-Host "`nTesting installation..." -ForegroundColor Yellow
& $BINARY_PATH --version

# Setup
Write-Host "`nRunning initial setup..." -ForegroundColor Yellow
& $BINARY_PATH update

Write-Host "`n✓ Installation complete!" -ForegroundColor Green
Write-Host "`nQuick start:" -ForegroundColor Cyan
Write-Host "  sniprun list              - Show available snips"
Write-Host "  sniprun add my-snip       - Create a custom snip"
Write-Host "  sniprun docker-clean      - Execute a snip"
Write-Host "`nFor more help: sniprun --help" -ForegroundColor Gray