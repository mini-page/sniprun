# Example Community Snips

Create these as `.yaml` files in the community repository:

## docker-clean.yaml
```yaml
name: docker-clean
description: Remove all Docker containers, images, volumes, and networks
command: docker system prune -a --volumes -f
args: []
category: docker
trust: community
```

## git-reset-hard.yaml
```yaml
name: git-reset-hard
description: Reset git branch to match remote, discarding all local changes
command: git fetch origin && git reset --hard origin/{{branch}}
args: [branch]
category: git
trust: community
```

## git-clean-merged.yaml
```yaml
name: git-clean-merged
description: Delete all local branches that have been merged
command: git branch --merged | grep -v '\*' | xargs -n 1 git branch -d
args: []
category: git
trust: community
```

## npm-fresh.yaml
```yaml
name: npm-fresh
description: Clean install npm dependencies
command: rm -rf node_modules package-lock.json && npm install
args: []
category: nodejs
trust: community
```

## port-kill.yaml
```yaml
name: port-kill
description: Kill process running on specified port
command: lsof -ti:{{port}} | xargs kill -9
args: [port]
category: system
trust: community
```

## disk-usage.yaml
```yaml
name: disk-usage
description: Show disk usage of current directory, sorted by size
command: du -sh * | sort -h
args: []
category: system
trust: community
```

## find-large.yaml
```yaml
name: find-large
description: Find files larger than specified size in current directory
command: find . -type f -size +{{size}}M -exec ls -lh {} \;
args: [size]
category: system
trust: community
```

## ssh-keygen-ed.yaml
```yaml
name: ssh-keygen-ed
description: Generate new ED25519 SSH key with email
command: ssh-keygen -t ed25519 -C "{{email}}"
args: [email]
category: security
trust: community
```

## ps-grep.yaml
```yaml
name: ps-grep
description: Search for running processes
command: ps aux | grep {{process}}
args: [process]
category: system
trust: community
```

## git-undo-commit.yaml
```yaml
name: git-undo-commit
description: Undo last commit but keep changes
command: git reset --soft HEAD~1
args: []
category: git
trust: community
```