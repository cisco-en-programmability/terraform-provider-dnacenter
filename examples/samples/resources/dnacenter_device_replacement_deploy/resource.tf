
terraform {
  required_providers {
    dnacenter = {
      version = "1.3.1-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_device_replacement_deploy" "example" {
  provider = dnacenter

  parameters {
    faulty_device_serial_number      = "string"
    replacement_device_serial_number = "string"
  }
}
