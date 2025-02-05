
data "dnacenter_dns_services" "example" {
  provider                 = dnacenter
  device_id                = "string"
  device_site_hierarchy_id = "string"
  device_site_id           = "string"
  end_time                 = 1609459200
  limit                    = 1
  offset                   = 1
  order                    = "string"
  server_ip                = "string"
  sort_by                  = "string"
  ssid                     = "string"
  start_time               = 1609459200
  xca_lle_rid              = "string"
}

output "dnacenter_dns_services_example" {
  value = data.dnacenter_dns_services.example.items
}
