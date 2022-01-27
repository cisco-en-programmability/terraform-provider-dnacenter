
data "dnacenter_platform_nodes_configuration_summary" "example" {
  provider = dnacenter
}

output "dnacenter_platform_nodes_configuration_summary_example" {
  value = data.dnacenter_platform_nodes_configuration_summary.example.item
}
