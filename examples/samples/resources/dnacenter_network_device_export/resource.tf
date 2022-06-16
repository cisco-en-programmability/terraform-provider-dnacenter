
terraform {
  required_providers {
    dnacenter = {
      version = "0.3.0"
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
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    device_uuids   = ["string"]
    id             = "string"
    operation_enum = "string"
    parameters     = ["string"]
    password       = "******"
  }
}