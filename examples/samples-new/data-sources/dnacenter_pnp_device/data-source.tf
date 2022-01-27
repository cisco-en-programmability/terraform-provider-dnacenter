terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

data "dnacenter_pnp_device" "example" {
  provider           = dnacenter
  cm_state           = ["string"]
  hostname           = "string"
  id                 = "618c2b7ff6b3b66f71acbaf8"
  last_contact       = "false"
  limit              = 1
  mac_address        = "string"
  name               = ["string"]
  offset             = 1
  onb_state          = ["string"]
  pid                = ["string"]
  project_id         = ["string"]
  project_name       = ["string"]
  serial_number      = ["string"]
  site_name          = "string"
  smart_account_id   = ["string"]
  sort               = ["string"]
  sort_order         = "string"
  source             = ["string"]
  state              = ["string"]
  virtual_account_id = ["string"]
  workflow_id        = ["string"]
  workflow_name      = ["string"]
}

output "dnacenter_pnp_device_example" {
  value = data.dnacenter_pnp_device.example.item_name
}
/*
data "dnacenter_pnp_device" "example" {
  provider = dnacenter
  id       = "618c2b7ff6b3b66f71acbaf8"
}

output "dnacenter_pnp_device_example" {
  value = data.dnacenter_pnp_device.example
}
*/