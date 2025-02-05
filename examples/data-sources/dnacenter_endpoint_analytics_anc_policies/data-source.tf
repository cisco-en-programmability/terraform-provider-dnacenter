
data "dnacenter_endpoint_analytics_anc_policies" "example" {
  provider = dnacenter
}

output "dnacenter_endpoint_analytics_anc_policies_example" {
  value = data.dnacenter_endpoint_analytics_anc_policies.example.items
}
