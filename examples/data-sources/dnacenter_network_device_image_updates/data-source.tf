
data "dnacenter_network_device_image_updates" "example" {
  provider           = dnacenter
  end_time           = 1609459200
  host_name          = "string"
  id                 = "string"
  image_name         = "string"
  limit              = 1
  management_address = "string"
  network_device_id  = "string"
  offset             = 1
  order              = "string"
  parent_id          = "string"
  sort_by            = "string"
  start_time         = 1609459200
  status             = "string"
}

output "dnacenter_network_device_image_updates_example" {
  value = data.dnacenter_network_device_image_updates.example.items
}
