
resource "dnacenter_interface_update" "example" {
  provider = dnacenter
  parameters {

    admin_status   = "string"
    description    = "string"
    interface_uuid = "string"
    vlan_id        = 1
    voice_vlan_id  = 1
  }
}

output "dnacenter_interface_update_example" {
  value = dnacenter_interface_update.example
}