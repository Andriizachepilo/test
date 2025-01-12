terraform {
  required_providers {
    cicd = {
      version = "0.1.0"
      source  = "local/cicd"
    }
  }
}

provider "cicd" {}



