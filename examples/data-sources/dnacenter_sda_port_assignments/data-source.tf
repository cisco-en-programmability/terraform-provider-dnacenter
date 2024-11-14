
data "dnacenter_sda_port_assignments" "example" {
  provider          = dnacenter
  data_vlan_name    = "string"
  fabric_id         = "string"
  interface_name    = "string"
  limit             = 1
  network_device_id = "string"
  offset            = 1
  voice_vlan_name   = "string"
}

output "dnacenter_sda_port_assignments_example" {
  value = data.dnacenter_sda_port_assignments.example.items
}
