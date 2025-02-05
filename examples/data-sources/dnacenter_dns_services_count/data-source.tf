
data "dnacenter_dns_services_count" "example" {
  provider                 = dnacenter
  device_id                = "string"
  device_site_hierarchy_id = "string"
  device_site_id           = "string"
  end_time                 = 1609459200
  server_ip                = "string"
  ssid                     = "string"
  start_time               = 1609459200
  xca_lle_rid              = "string"
}

output "dnacenter_dns_services_count_example" {
  value = data.dnacenter_dns_services_count.example.item
}
