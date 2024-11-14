
data "dnacenter_sda_multicast_virtual_networks_count" "example" {
  provider  = dnacenter
  fabric_id = "string"
}

output "dnacenter_sda_multicast_virtual_networks_count_example" {
  value = data.dnacenter_sda_multicast_virtual_networks_count.example.item
}
