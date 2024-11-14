
data "dnacenter_interfaces_count" "example" {
  provider                   = dnacenter
  end_time                   = 1609459200
  interface_id               = "string"
  interface_name             = "string"
  network_device_id          = "string"
  network_device_ip_address  = "string"
  network_device_mac_address = "string"
  site_hierarchy             = "string"
  site_hierarchy_id          = "string"
  site_id                    = "string"
  start_time                 = 1609459200
}

output "dnacenter_interfaces_count_example" {
  value = data.dnacenter_interfaces_count.example.item
}
