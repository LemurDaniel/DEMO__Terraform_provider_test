# Terraform Provider Demo Test

Testing terraform providers for maybe future implementation

```powershell
cd terraform_test

# On changes
./build.ps1

terraform init -plugin-dir="./registry.terraform.io/" -reconfigure

terraform apply
```
