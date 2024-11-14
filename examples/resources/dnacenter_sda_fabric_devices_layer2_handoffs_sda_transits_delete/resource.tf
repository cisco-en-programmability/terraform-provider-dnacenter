
resource "dnacenter_sda_fabric_devices_layer2_handoffs_sda_transits_delete" "example" {
  provider          = dnacenter
  fabric_id         = "string"
  network_device_id = "string"
}

output "dnacenter_sda_fabric_devices_layer2_handoffs_sda_transits_delete_example" {
  value = dnacenter_sda_fabric_devices_layer2_handoffs_sda_transits_delete.example
}