
data "dnacenter_network_device_replacements_id" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_network_device_replacements_id_example" {
  value = data.dnacenter_network_device_replacements_id.example.item
}
