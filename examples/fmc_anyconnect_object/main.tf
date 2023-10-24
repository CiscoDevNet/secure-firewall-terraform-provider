terraform {
  required_providers {
    fmc = {
      source = "CiscoDevNet/fmc"
      # version = "0.2.0"
    }
  }
}

provider "fmc" {
  fmc_username = var.fmc_username
  fmc_password = var.fmc_password
  fmc_host = var.fmc_host
  fmc_insecure_skip_verify = var.fmc_insecure_skip_verify
}

data "fmc_anyconnect_package" "obj" {
    name = "anyconnect-pkg"
}

output "anyconnect_package_output" {
    value = data.fmc_anyconnect_package.obj
}
