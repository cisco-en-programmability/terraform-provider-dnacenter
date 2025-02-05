
data "dnacenter_network_bugs_results_bugs_id" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_network_bugs_results_bugs_id_example" {
  value = data.dnacenter_network_bugs_results_bugs_id.example.item
}
