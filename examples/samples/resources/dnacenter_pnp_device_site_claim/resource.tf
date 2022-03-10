
terraform {
  required_providers {
    dnacenter = {
      version = "0.2.0-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_pnp_device_site_claim" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    device_id = "61fd411fd86a6c29631156f2"
    site_id   = "9625864c-2228-42ec-ac9c-1d0e8e099825"
    type      = "Default"
    hostname  = "Catalyst2"
    image_info {
      image_id = "6af2b040-a312-4f57-8c8e-21f5e3e07597"
      skip     = "true"
    }
    config_info {
      config_id = "60657208-dc2c-433f-be40-6c002d19d0fb"
      config_parameters {
        key   = ""
        value = ""
      }
    }
  }
}

output "dnacenter_pnp_device_site_claim_example" {
  value = dnacenter_pnp_device_site_claim.example
}