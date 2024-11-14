
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.33-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_pnp_device" "example" {
  provider = dnacenter
  # lifecycle {
  #   create_before_destroy = true
  # }
  parameters {

    # id = "61f1c861264f342e4fa1a78e"
    device_info {

      serial_number = "FLM2213W05S"
      stack         = "false"
      # state= "Unclaimed"



      sudi_required = "false"
      hostname      = "FLM2213W05W"


    }
  }
}

output "dnacenter_pnp_device_example" {
  value = dnacenter_pnp_device.example
}
