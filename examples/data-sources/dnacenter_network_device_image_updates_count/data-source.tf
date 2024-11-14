
data "dnacenter_network_device_image_updates_count" "example" {
  provider           = dnacenter
  end_time           = 1609459200
  host_name          = "string"
  id                 = "string"
  image_name         = "string"
  management_address = "string"
  network_device_id  = "string"
  parent_id          = "string"
  start_time         = 1609459200
  status             = "string"
}

output "dnacenter_network_device_image_updates_count_example" {
  value = data.dnacenter_network_device_image_updates_count.example.item
}
