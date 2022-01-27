
data "dnacenter_network_device_count" "example" {
  provider  = dnacenter
  device_id = "string"
}

output "dnacenter_network_device_count_example" {
  value = data.dnacenter_network_device_count.example.item_name
}

data "dnacenter_network_device_count" "example" {
  provider  = dnacenter
  device_id = "string"
}

output "dnacenter_network_device_count_example" {
  value = data.dnacenter_network_device_count.example.item_id
}
