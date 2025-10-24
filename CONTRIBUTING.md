# Contributing to sniprun

We welcome contributions to `sniprun`! To make the process as smooth as possible, please follow these guidelines.

## How to Contribute

### 1. Fork the Repository

First, fork the `sniprun` repository to your GitHub account.

### 2. Clone Your Fork

Clone your forked repository to your local machine:

```bash
git clone https://github.com/mini-page/sniprun.git
cd sniprun
```

### 3. Create a New Branch

Create a new branch for your feature or bug fix:

```bash
git checkout -b feature/your-feature-name
# or
git checkout -b bugfix/your-bug-fix-name
```

### 4. Make Your Changes

- Implement your feature or bug fix.
- Ensure your code adheres to the existing coding style and conventions.
- Write tests for your changes to ensure they work as expected and prevent regressions.

### 5. Test Your Changes

Run the test suite to make sure everything is still working correctly:

```bash
go test ./...
```

### 6. Commit Your Changes

Commit your changes with a clear and concise commit message. Follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification (e.g., `feat: add new feature`, `fix: resolve bug`).

```bash
git commit -m "feat: Add a new awesome feature"
```

### 7. Push to Your Fork

Push your changes to your forked repository:

```bash
git push origin feature/your-feature-name
```

### 8. Create a Pull Request

Go to the original `sniprun` repository on GitHub and open a new Pull Request from your forked branch to the `main` branch. Provide a detailed description of your changes.

## Code Style

- We follow standard Go formatting practices. Please run `go fmt ./...` before committing.
- Ensure your code is well-commented where necessary.

## Reporting Bugs

If you find a bug, please open an issue on GitHub. Provide a clear description of the bug, steps to reproduce it, and expected behavior.

## Feature Requests

If you have an idea for a new feature, please open an issue on GitHub to discuss it. This helps ensure that your efforts align with the project's goals.
