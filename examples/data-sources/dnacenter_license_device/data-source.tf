
data "dnacenter_license_device" "example" {
  provider = dnacenter
}

output "dnacenter_license_device_example" {
  value = data.dnacenter_license_device.example.items
}
