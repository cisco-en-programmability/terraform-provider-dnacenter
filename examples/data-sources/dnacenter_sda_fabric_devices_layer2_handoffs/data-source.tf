
data "dnacenter_sda_fabric_devices_layer2_handoffs" "example" {
  provider          = dnacenter
  fabric_id         = "string"
  limit             = 1
  network_device_id = "string"
  offset            = 1
}

output "dnacenter_sda_fabric_devices_layer2_handoffs_example" {
  value = data.dnacenter_sda_fabric_devices_layer2_handoffs.example.items
}
