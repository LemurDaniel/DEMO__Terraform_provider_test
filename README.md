# Terraform Provider Demo Test

Run the following command to build the provider

```powershell
.\build.ps1
```

## Test sample configuration


```powershell
cd terraform_test
terraform init -plugin-dir="./registry.terraform.io/"
terraform apply
```
