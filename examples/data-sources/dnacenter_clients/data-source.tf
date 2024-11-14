
data "dnacenter_clients" "example" {
  provider                      = dnacenter
  attribute                     = "string"
  band                          = "string"
  connected_network_device_name = "string"
  end_time                      = 1609459200
  ipv4_address                  = "string"
  ipv6_address                  = "string"
  limit                         = 1
  mac_address                   = "string"
  offset                        = 1
  order                         = "string"
  os_type                       = "string"
  os_version                    = "string"
  site_hierarchy                = "string"
  site_hierarchy_id             = "string"
  site_id                       = "string"
  sort_by                       = "string"
  ssid                          = "string"
  start_time                    = 1609459200
  type                          = "string"
  view                          = "string"
  wlc_name                      = "string"
  xca_lle_rid                   = "string"
}

output "dnacenter_clients_example" {
  value = data.dnacenter_clients.example.items
}

data "dnacenter_clients" "example" {
  provider    = dnacenter
  attribute   = "string"
  end_time    = 1609459200
  id          = "string"
  start_time  = 1609459200
  view        = "string"
  xca_lle_rid = "string"
}

output "dnacenter_clients_example" {
  value = data.dnacenter_clients.example.item
}
