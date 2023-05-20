
$HOSTNAME = "terraform.local"
$NAMESPACE = "lemurdaniel"
$NAME = "azurepim"
$VERSION = "0.0.1"
$TARGET = "windows_amd64"


$installPath = "./terraform_test/registry.terraform.io/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${TARGET}"
$binaryName = "terraform-provider-$($NAME)_v$($VERSION)_windows_amd64"

New-Item -ItemType Directory -Path $installPath -ErrorAction SilentlyContinue
go build -o "$installPath/$binaryName"