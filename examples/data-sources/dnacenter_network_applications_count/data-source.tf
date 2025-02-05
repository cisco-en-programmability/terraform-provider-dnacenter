
data "dnacenter_network_applications_count" "example" {
  provider           = dnacenter
  application_name   = "string"
  business_relevance = "string"
  end_time           = 1609459200
  site_id            = "string"
  ssid               = "string"
  start_time         = 1609459200
  xca_lle_rid        = "string"
}

output "dnacenter_network_applications_count_example" {
  value = data.dnacenter_network_applications_count.example.item
}
