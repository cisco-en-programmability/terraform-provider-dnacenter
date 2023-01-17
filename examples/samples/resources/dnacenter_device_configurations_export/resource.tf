
terraform {
  required_providers {
    dnacenter = {
      version = "1.0.16-beta"
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
    device_id = ["string"]
    password  = "******"
  }
}