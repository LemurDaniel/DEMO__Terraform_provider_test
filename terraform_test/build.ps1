
$HOSTNAME = "terraform.local"
$NAMESPACE = "lemurdaniel"
$NAME = "todotest"
$VERSION = "0.0.1"
$TARGET = "windows_amd64"

$toplevel = git rev-parse --show-toplevel
$current = $PSScriptRoot

$null = Set-Location -Path $toplevel
Write-Host "Preparing..."
go mod tidy
go mod vendor


$installPath = "$toplevel/terraform_test/registry.terraform.io/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${TARGET}"
$binaryName = "terraform-provider-$($NAME)_v$($VERSION)_windows_amd64"

Remove-Item -Path "$toplevel/terraform_test/.terraform.lock.hcl" -ErrorAction SilentlyContinue
Remove-Item -Path "$installPath/$binaryName" -ErrorAction SilentlyContinue
New-Item -ItemType Directory -Path $installPath -ErrorAction SilentlyContinue

Write-Host "Building..."
go build -o "$installPath/$binaryName"

$null = Set-Location -Path $current
terraform init -plugin-dir="./registry.terraform.io/" -reconfigure

$env:TF_LOG = "INFO"