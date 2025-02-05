
data "dnacenter_network_bugs_results_bugs" "example" {
  provider     = dnacenter
  device_count = 1.0
  id           = "string"
  limit        = 1
  offset       = 1
  order        = "string"
  severity     = "string"
  sort_by      = "string"
}

output "dnacenter_network_bugs_results_bugs_example" {
  value = data.dnacenter_network_bugs_results_bugs.example.items
}
