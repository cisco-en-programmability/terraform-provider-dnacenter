
data "dnacenter_aaa_services" "example" {
  provider                 = dnacenter
  device_id                = "string"
  device_name              = "string"
  device_site_hierarchy_id = "string"
  end_time                 = 1609459200
  limit                    = 1
  offset                   = 1
  order                    = "string"
  server_ip                = "string"
  site_hierarchy           = "string"
  site_id                  = "string"
  sort_by                  = "string"
  start_time               = 1609459200
  xca_lle_rid              = "string"
}

output "dnacenter_aaa_services_example" {
  value = data.dnacenter_aaa_services.example.items
}
