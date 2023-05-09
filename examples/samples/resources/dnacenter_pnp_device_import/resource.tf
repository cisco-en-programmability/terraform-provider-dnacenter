
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.4-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_pnp_device_import" "example" {
  provider = dnacenter

  parameters {
    payload {
      device_info {
        serial_number = "FOCTEST5AB"
      }
    }
    #   payload {
    #     device_info {
    #       serial_number = "FOCTEST511"
    #     }
    #   }
  }
}

output "dnacenter_pnp_device_import_example" {
  value = dnacenter_pnp_device_import.example
}
