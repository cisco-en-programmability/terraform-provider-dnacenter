
data "dnacenter_clients_count" "example" {
  provider                      = dnacenter
  band                          = "string"
  connected_network_device_name = "string"
  end_time                      = 1609459200
  ipv4_address                  = "string"
  ipv6_address                  = "string"
  mac_address                   = "string"
  os_type                       = "string"
  os_version                    = "string"
  site_hierarchy                = "string"
  site_hierarchy_id             = "string"
  site_id                       = "string"
  ssid                          = "string"
  start_time                    = 1609459200
  type                          = "string"
  wlc_name                      = "string"
  xca_lle_rid                   = "string"
}

output "dnacenter_clients_count_example" {
  value = data.dnacenter_clients_count.example.item
}
