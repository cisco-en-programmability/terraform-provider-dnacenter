
resource "dnacenter_wireless_dynamic_interface" "example" {
  provider = dnacenter
  parameters {

    interface_name = "string"
    vlan_id        = 1.0
  }
}

output "dnacenter_wireless_dynamic_interface_example" {
  value = dnacenter_wireless_dynamic_interface.example
}