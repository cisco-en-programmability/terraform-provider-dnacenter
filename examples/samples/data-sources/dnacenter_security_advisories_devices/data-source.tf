
data "dnacenter_security_advisories_devices" "example" {
  provider    = dnacenter
  advisory_id = "string"
}

output "dnacenter_security_advisories_devices_example" {
  value = data.dnacenter_security_advisories_devices.example.items
}
