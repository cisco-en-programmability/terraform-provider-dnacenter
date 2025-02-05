
data "dnacenter_endpoint_analytics_dictionaries" "example" {
  provider           = dnacenter
  include_attributes = "false"
}

output "dnacenter_endpoint_analytics_dictionaries_example" {
  value = data.dnacenter_endpoint_analytics_dictionaries.example.items
}
