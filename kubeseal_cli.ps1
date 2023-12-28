# Create the "C:\bin\kubeseal" directory if it doesn't exist
Write-Host "Creating the 'C:\bin\kubeseal' directory if it doesn't exist..."
$binDir = "C:\bin\kubeseal"
New-Item -ItemType Directory -Path $binDir -ErrorAction Ignore

# Get the latest release information from the GitHub API
Write-Host "Retrieving the latest release information from the GitHub API..."
$releaseInfo = Invoke-WebRequest -Uri https://api.github.com/repos/bitnami-labs/sealed-secrets/releases/latest

# Extract the download URL for the Windows kubeseal binary
Write-Host "Extracting the download URL for the Windows kubeseal binary..."
$downloadUrl = ($releaseInfo | ConvertFrom-Json).assets | Where-Object { $_.name -match "^kubeseal-.*-windows-amd64\.tar\.gz$" } | Select-Object -ExpandProperty "browser_download_url"

# Download the kubeseal binary
Write-Host "Downloading the kubeseal binary..."
$tempFile = [System.IO.Path]::GetTempFileName()
Invoke-WebRequest -Uri $downloadUrl -OutFile $tempFile

# Uncompress the .tar.gz file to the installation directory
Write-Host "Uncompressing the .tar.gz file to $binDir..."
tar -xvzf $tempFile -C $binDir

# Add the installation directory to the PATH environment variable
Write-Host "Adding the '$binDir' directory to the PATH environment variable..."
[Environment]::SetEnvironmentVariable("PATH", "$binDir;$env:PATH", "Machine")

# Refresh the PATH
Write-Host "Refreshing the PATH..."
$env:Path = [System.Environment]::GetEnvironmentVariable("Path", "Machine")

# Test that kubeseal is installed and working
Write-Host "Testing that kubeseal is installed and working..."
$kubesealOutput = & "$binDir\kubeseal" --version
Write-Output $kubesealOutput