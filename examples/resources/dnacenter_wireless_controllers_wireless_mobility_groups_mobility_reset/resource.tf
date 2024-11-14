
resource "dnacenter_wireless_controllers_wireless_mobility_groups_mobility_reset" "example" {
  provider = dnacenter
  parameters {

    network_device_id = "string"
  }
}

output "dnacenter_wireless_controllers_wireless_mobility_groups_mobility_reset_example" {
  value = dnacenter_wireless_controllers_wireless_mobility_groups_mobility_reset.example
}