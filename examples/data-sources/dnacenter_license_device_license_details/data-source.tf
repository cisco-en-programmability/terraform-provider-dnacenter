
data "dnacenter_license_device_license_details" "example" {
  provider    = dnacenter
  device_uuid = "string"
}

output "dnacenter_license_device_license_details_example" {
  value = data.dnacenter_license_device_license_details.example.items
}
