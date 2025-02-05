
data "dnacenter_network_bugs_results_bugs_count" "example" {
  provider     = dnacenter
  device_count = 1.0
  id           = "string"
  severity     = "string"
}

output "dnacenter_network_bugs_results_bugs_count_example" {
  value = data.dnacenter_network_bugs_results_bugs_count.example.item
}
