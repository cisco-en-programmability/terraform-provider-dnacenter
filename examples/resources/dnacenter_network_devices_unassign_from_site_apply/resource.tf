
resource "dnacenter_network_devices_unassign_from_site_apply" "example" {
  provider = dnacenter
  parameters {

    device_ids = ["string"]
  }
}

output "dnacenter_network_devices_unassign_from_site_apply_example" {
  value = dnacenter_network_devices_unassign_from_site_apply.example
}