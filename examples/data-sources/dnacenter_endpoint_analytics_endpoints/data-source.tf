
data "dnacenter_endpoint_analytics_endpoints" "example" {
  provider = dnacenter
  ep_id    = "string"
  include  = "string"
}

output "dnacenter_endpoint_analytics_endpoints_example" {
  value = data.dnacenter_endpoint_analytics_endpoints.example.item
}
