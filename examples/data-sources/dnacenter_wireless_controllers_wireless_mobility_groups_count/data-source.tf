
data "dnacenter_wireless_controllers_wireless_mobility_groups_count" "example" {
  provider = dnacenter
}

output "dnacenter_wireless_controllers_wireless_mobility_groups_count_example" {
  value = data.dnacenter_wireless_controllers_wireless_mobility_groups_count.example.item
}
