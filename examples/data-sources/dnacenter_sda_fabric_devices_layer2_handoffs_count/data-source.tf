
data "dnacenter_sda_fabric_devices_layer2_handoffs_count" "example" {
  provider          = dnacenter
  fabric_id         = "string"
  network_device_id = "string"
}

output "dnacenter_sda_fabric_devices_layer2_handoffs_count_example" {
  value = data.dnacenter_sda_fabric_devices_layer2_handoffs_count.example.item
}
