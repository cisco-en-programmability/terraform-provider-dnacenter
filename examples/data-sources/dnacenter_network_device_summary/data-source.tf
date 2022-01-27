
data "dnacenter_network_device_summary" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_network_device_summary_example" {
  value = data.dnacenter_network_device_summary.example.item
}
