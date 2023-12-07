
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.31-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_network_device_export" "example" {
  provider = dnacenter

  parameters {
    device_uuids   = ["string"]
    id             = "string"
    operation_enum = "string"
    parameters     = ["string"]
    password       = "******"
  }
}
