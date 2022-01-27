
data "dnacenter_device_family_identifiers_details" "example" {
  provider = dnacenter
}

output "dnacenter_device_family_identifiers_details_example" {
  value = data.dnacenter_device_family_identifiers_details.example.items
}
