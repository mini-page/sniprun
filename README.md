![sniprun logo](./assets/logo.png)

# sniprun

**Run complex commands with short, memorable snips.**

![Gemini API Badge](https://img.shields.io/badge/Powered%20by-Gemini%20API-blue?style=flat-square)

`sniprun` is a community-driven CLI tool that lets you store and execute simplified aliases for complex commands. Think of it as a combination of `tldr` and executable aliases, but shareable and version-controlled.

## ✨ Features

- 🚀 **Execute complex commands** with simple aliases
- 🌐 **Community-contributed** snip repository
- 🔒 **Security validation** using Gemini API (optional)
- 📝 **Argument interpolation** for flexible commands
- 🏠 **Local custom snips** alongside community ones
- 💻 **Cross-platform** (Windows, macOS, Linux)

## ⚡ Quick Start

```bash
# Install
go install github.com/mini-page/sniprun@latest

# Fetch community snips
sniprun update

# List available snips
sniprun list

# Execute a snip
sniprun docker-clean

# Execute with arguments
sniprun git-reset-hard main
```

## 🛠️ Installation

### From Source

```bash
git clone https://github.com/mini-page/sniprun
cd sniprun
go build -o sniprun
sudo mv sniprun /usr/local/bin/
```

### Set up Gemini API (Optional)

For security validation:

```bash
export GEMINI_API_KEY="your-api-key"
```

Without an API key, security checks are skipped.

## 🚀 Usage

### Basic Commands

```bash
# Update community snips
sniprun update

# List all available snips
sniprun list

# List by category
sniprun list --category git

# Show what a snip does
sniprun explain docker-clean

# Execute a snip
sniprun docker-clean

# Execute with arguments
sniprun port-kill 3000
```

### Creating Custom Snips

```bash
# Interactive creation
sniprun add my-backup

# You'll be prompted for:
# - Description
# - Category
# - Command
# - Arguments (optional)
```

Example snip with arguments:

```yaml
name: git-commit-push
description: Commit and push with message
command: git add . && git commit -m "{{message}}" && git push
args: [message]
category: git
trust: local
```

Usage: `sniprun git-commit-push "fix: bug fix"`

### Advanced Usage

```bash
# Skip security checks
sniprun docker-clean --skip-check

# Output command for shell evaluation (for cd, export, etc.)
eval $(sniprun my-cd-command --source)

# Remove a local snip
sniprun remove my-snip

# Force remove without confirmation
sniprun remove my-snip --force
```

## 🔒 Security

`sniprun` includes optional security validation using Google's Gemini API:

- ✅ **Safe**: Normal commands execute without prompts
- ⚠️ **Warning**: Destructive commands (rm, format) prompt for confirmation
- ❌ **Dangerous**: Malicious commands are blocked

Security levels:
- 🔧 **Local**: Your custom snips
- 🌐 **Community**: Public repository snips
- ✓ **Verified**: Reviewed and approved snips

## 📂 File Structure

```
~/.sniprun/
├── snips/
│   ├── local/          # Your custom snips
│   │   └── my-snip.yaml
│   └── community/      # Downloaded community snips
│       └── docker-clean.yaml
```

## 📝 Snip Format

```yaml
name: example-snip
description: What this snip does
command: echo "Hello {{name}}"
args: [name]              # Optional
category: examples
trust: local              # local | community | verified
```

## 🤝 Contributing

We welcome community contributions!

### Adding Snips

1. Fork the [sniprun/snips](https://github.com/sniprun/snips) repository
2. Create a new `.yaml` file in the appropriate category folder
3. Test your snip locally
4. Submit a pull request

### Guidelines

- Keep commands focused and reusable
- Add clear descriptions
- Use meaningful argument names
- Test on multiple platforms when possible
- Avoid platform-specific commands unless necessary

## 🆚 Comparison to Other Tools

| Tool | Purpose | sniprun Advantage |
|------|---------|-------------------|
| `tldr` | Command examples | Actually executes |
| `alias` | Shell aliases | Shareable, versioned, cross-shell |
| `just`/`task` | Project automation | Global, not project-specific |
| `cheat.sh` | Command cheatsheet | Integrated execution |

## 🗺️ Roadmap

- [ ] Web UI for browsing snips
- [ ] Snip ratings and reviews
- [ ] Multi-command workflows
- [ ] Conditional execution
- [ ] Platform-specific snips
- [ ] Package manager integration (brew, apt, choco)

## 📄 License

MIT License - see LICENSE file

## ❓ Support

- Issues: https://github.com/mini-page/sniprun/issues
- Discussions: https://github.com/mini-page/sniprun/discussions
- Community snips: https://github.com/sniprun/snips
