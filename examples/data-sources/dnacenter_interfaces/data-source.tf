
data "dnacenter_interfaces" "example" {
  provider                   = dnacenter
  attribute                  = "string"
  end_time                   = 1609459200
  interface_id               = "string"
  interface_name             = "string"
  limit                      = 1
  network_device_id          = "string"
  network_device_ip_address  = "string"
  network_device_mac_address = "string"
  offset                     = 1
  order                      = "string"
  site_hierarchy             = "string"
  site_hierarchy_id          = "string"
  site_id                    = "string"
  sort_by                    = "string"
  start_time                 = 1609459200
  view                       = "string"
}

output "dnacenter_interfaces_example" {
  value = data.dnacenter_interfaces.example.items
}

data "dnacenter_interfaces" "example" {
  provider   = dnacenter
  attribute  = "string"
  end_time   = 1609459200
  id         = "string"
  start_time = 1609459200
  view       = "string"
}

output "dnacenter_interfaces_example" {
  value = data.dnacenter_interfaces.example.item
}
