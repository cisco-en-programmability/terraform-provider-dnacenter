
data "dnacenter_sda_port_assignments_count" "example" {
  provider          = dnacenter
  data_vlan_name    = "string"
  fabric_id         = "string"
  interface_name    = "string"
  network_device_id = "string"
  voice_vlan_name   = "string"
}

output "dnacenter_sda_port_assignments_count_example" {
  value = data.dnacenter_sda_port_assignments_count.example.item
}
