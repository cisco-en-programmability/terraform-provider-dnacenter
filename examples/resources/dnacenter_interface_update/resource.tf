
resource "dnacenter_interface_update" "example" {
  provider        = dnacenter
  deployment_mode = "string"
  interface_uuid  = "string"
  parameters {

    admin_status  = "string"
    description   = "string"
    vlan_id       = 1
    voice_vlan_id = 1
  }
}

output "dnacenter_interface_update_example" {
  value = dnacenter_interface_update.example
}