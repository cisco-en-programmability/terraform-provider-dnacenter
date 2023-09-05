
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.13-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_device_configurations_export" "example" {
  provider = dnacenter

  parameters {
    device_id = ["3923aed0-16e5-4ed0-b430-ff6dcfd9c517"]
    password  = "Hola123*"
  }
}
