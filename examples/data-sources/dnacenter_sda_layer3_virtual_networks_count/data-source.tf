
data "dnacenter_sda_layer3_virtual_networks_count" "example" {
  provider         = dnacenter
  anchored_site_id = "string"
  fabric_id        = "string"
}

output "dnacenter_sda_layer3_virtual_networks_count_example" {
  value = data.dnacenter_sda_layer3_virtual_networks_count.example.item
}
