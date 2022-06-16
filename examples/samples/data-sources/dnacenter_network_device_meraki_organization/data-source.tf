
data "dnacenter_network_device_meraki_organization" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_network_device_meraki_organization_example" {
  value = data.dnacenter_network_device_meraki_organization.example.items
}
