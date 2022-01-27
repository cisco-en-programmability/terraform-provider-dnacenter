
data "dnacenter_pnp_device" "example" {
  provider           = dnacenter
  cm_state           = ["string"]
  hostname           = "string"
  id                 = "string"
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

data "dnacenter_pnp_device" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_pnp_device_example" {
  value = data.dnacenter_pnp_device.example.item_id
}
