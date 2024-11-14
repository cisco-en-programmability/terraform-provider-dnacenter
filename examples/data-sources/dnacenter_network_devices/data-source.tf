
data "dnacenter_network_devices" "example" {
  provider              = dnacenter
  attribute             = "string"
  end_time              = 1609459200
  family                = "string"
  health_score          = "string"
  id                    = "string"
  limit                 = 1
  mac_address           = "string"
  maintenance_mode      = "false"
  management_ip_address = "string"
  offset                = 1
  order                 = "string"
  role                  = "string"
  serial_number         = "string"
  site_hierarchy        = "string"
  site_hierarchy_id     = "string"
  site_id               = "string"
  software_version      = "string"
  sort_by               = "string"
  start_time            = 1609459200
  type                  = "string"
  view                  = "string"
}

output "dnacenter_network_devices_example" {
  value = data.dnacenter_network_devices.example.items
}

data "dnacenter_network_devices" "example" {
  provider   = dnacenter
  attribute  = "string"
  end_time   = 1609459200
  id         = "string"
  start_time = 1609459200
  view       = "string"
}

output "dnacenter_network_devices_example" {
  value = data.dnacenter_network_devices.example.item
}
