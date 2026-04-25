# install.ps1

$arch = if ([System.Environment]::Is64BitOperatingSystem) { "amd64" } else { "386" }
$url = "https://github.com/insigmo/jetreset/releases/latest/download/jetreset-windows-$arch.exe"
Invoke-WebRequest -Uri $url -OutFile "jetreset.exe"
Write-Host "Run exec file for reseting: .\jetreset.exe"
