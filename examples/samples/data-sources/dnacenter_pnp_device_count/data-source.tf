
data "dnacenter_pnp_device_count" "example" {
  provider           = dnacenter
  cm_state           = ["string"]
  last_contact       = "false"
  name               = ["string"]
  onb_state          = ["string"]
  pid                = ["string"]
  project_id         = ["string"]
  project_name       = ["string"]
  serial_number      = ["string"]
  smart_account_id   = ["string"]
  source             = ["string"]
  state              = ["string"]
  virtual_account_id = ["string"]
  workflow_id        = ["string"]
  workflow_name      = ["string"]
}

output "dnacenter_pnp_device_count_example" {
  value = data.dnacenter_pnp_device_count.example.item
}
