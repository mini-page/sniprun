# PowerShell-Specific Snip Examples

These snips are designed specifically for PowerShell/Windows environments:

## ps-admin.yaml
```yaml
name: ps-admin
description: Start a new PowerShell session as Administrator
command: Start-Process powershell -Verb RunAs
args: []
category: powershell
trust: community
```

## win-clean-temp.yaml
```yaml
name: win-clean-temp
description: Clean Windows temporary files
command: Remove-Item -Path $env:TEMP\* -Recurse -Force -ErrorAction SilentlyContinue
args: []
category: windows
trust: community
```

## ps-history-clear.yaml
```yaml
name: ps-history-clear
description: Clear PowerShell command history
command: Clear-History; Remove-Item (Get-PSReadlineOption).HistorySavePath -ErrorAction SilentlyContinue
args: []
category: powershell
trust: community
```

## win-network-reset.yaml
```yaml
name: win-network-reset
description: Reset network adapters
command: Get-NetAdapter | Restart-NetAdapter
args: []
category: windows
trust: community
```

## ps-profile-edit.yaml
```yaml
name: ps-profile-edit
description: Open PowerShell profile in default editor
command: notepad $PROFILE
args: []
category: powershell
trust: community
```

## win-processes-cpu.yaml
```yaml
name: win-processes-cpu
description: Show top CPU consuming processes
command: Get-Process | Sort-Object CPU -Descending | Select-Object -First 10 Name, CPU, PM
args: []
category: windows
trust: community
```

## ps-module-update.yaml
```yaml
name: ps-module-update
description: Update all PowerShell modules
command: Get-InstalledModule | Update-Module -Force
args: []
category: powershell
trust: community
```

## win-firewall-status.yaml
```yaml
name: win-firewall-status
description: Check Windows Firewall status for all profiles
command: Get-NetFirewallProfile | Select-Object Name, Enabled
args: []
category: windows
trust: community
```

## ps-env-path.yaml
```yaml
name: ps-env-path
description: Display PATH environment variable in readable format
command: $env:Path -split ';' | ForEach-Object { $_ }
args: []
category: powershell
trust: community
```

## win-service-restart.yaml
```yaml
name: win-service-restart
description: Restart a Windows service by name
command: Restart-Service -Name {{service}} -Force
args: [service]
category: windows
trust: community
```

## ps-json-pretty.yaml
```yaml
name: ps-json-pretty
description: Pretty print JSON from clipboard
command: Get-Clipboard | ConvertFrom-Json | ConvertTo-Json -Depth 10
args: []
category: powershell
trust: community
```

## win-wifi-password.yaml
```yaml
name: win-wifi-password
description: Show saved WiFi password for a network
command: netsh wlan show profile name="{{ssid}}" key=clear
args: [ssid]
category: windows
trust: community
```

## ps-error-last.yaml
```yaml
name: ps-error-last
description: Show details of last error
command: $Error[0] | Format-List * -Force
args: []
category: powershell
trust: community
```

## win-disk-cleanup.yaml
```yaml
name: win-disk-cleanup
description: Run Windows Disk Cleanup utility
command: Start-Process cleanmgr -ArgumentList '/sagerun:1' -Verb RunAs
args: []
category: windows
trust: community
```

## ps-execution-policy.yaml
```yaml
name: ps-execution-policy
description: Set PowerShell execution policy to RemoteSigned for current user
command: Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
args: []
category: powershell
trust: community
```

## Usage Examples

```powershell
# Clean temp files
sniprun win-clean-temp

# Restart IIS service
sniprun win-service-restart w3svc

# Get WiFi password
sniprun win-wifi-password "MyNetwork"

# Start admin PowerShell
sniprun ps-admin

# Update all PowerShell modules
sniprun ps-module-update
```