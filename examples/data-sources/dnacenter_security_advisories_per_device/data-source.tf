
data "dnacenter_security_advisories_per_device" "example" {
  provider  = dnacenter
  device_id = "string"
}

output "dnacenter_security_advisories_per_device_example" {
  value = data.dnacenter_security_advisories_per_device.example.item
}
