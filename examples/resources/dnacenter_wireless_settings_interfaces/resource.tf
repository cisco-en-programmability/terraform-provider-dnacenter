
resource "dnacenter_wireless_settings_interfaces" "example" {
  provider = dnacenter

  parameters {

    id             = "string"
    interface_name = "string"
    vlan_id        = 1
  }
}

output "dnacenter_wireless_settings_interfaces_example" {
  value = dnacenter_wireless_settings_interfaces.example
}