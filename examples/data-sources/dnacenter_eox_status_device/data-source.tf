
data "dnacenter_eox_status_device" "example" {
  provider = dnacenter
}

output "dnacenter_eox_status_device_example" {
  value = data.dnacenter_eox_status_device.example.items
}

data "dnacenter_eox_status_device" "example" {
  provider  = dnacenter
  device_id = "string"
}

output "dnacenter_eox_status_device_example" {
  value = data.dnacenter_eox_status_device.example.item
}
