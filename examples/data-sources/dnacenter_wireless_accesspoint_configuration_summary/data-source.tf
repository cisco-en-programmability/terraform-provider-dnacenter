
data "dnacenter_wireless_accesspoint_configuration_summary" "example" {
  provider = dnacenter
  key      = "string"
}

output "dnacenter_wireless_accesspoint_configuration_summary_example" {
  value = data.dnacenter_wireless_accesspoint_configuration_summary.example.item
}
