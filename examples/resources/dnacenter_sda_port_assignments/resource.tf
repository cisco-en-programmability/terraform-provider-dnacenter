
resource "dnacenter_sda_port_assignments" "example" {
  provider = dnacenter

  parameters {

    authenticate_template_name = "string"
    connected_device_type      = "string"
    data_vlan_name             = "string"
    fabric_id                  = "string"
    id                         = "string"
    interface_description      = "string"
    interface_name             = "string"
    network_device_id          = "string"
    scalable_group_name        = "string"
    security_group_name        = "string"
    voice_vlan_name            = "string"
  }
}

output "dnacenter_sda_port_assignments_example" {
  value = dnacenter_sda_port_assignments.example
}