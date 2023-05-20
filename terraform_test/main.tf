terraform {

  required_providers {

    hashicups = {
      version = "0.0.1"
      source  = "terraform.local/lemurdaniel/azurepim"
    }
  }

  backend "local" {

  }
  
}

