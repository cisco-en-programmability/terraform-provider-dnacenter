
resource "dnacenter_sda_fabric_devices_layer2_handoffs" "example" {
  provider = dnacenter

  parameters {

    external_vlan_id  = 1
    fabric_id         = "string"
    id                = "string"
    interface_name    = "string"
    internal_vlan_id  = 1
    network_device_id = "string"
  }
}

output "dnacenter_sda_fabric_devices_layer2_handoffs_example" {
  value = dnacenter_sda_fabric_devices_layer2_handoffs.example
}