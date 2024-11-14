
data "dnacenter_sda_transit_networks_count" "example" {
  provider = dnacenter
  type     = "string"
}

output "dnacenter_sda_transit_networks_count_example" {
  value = data.dnacenter_sda_transit_networks_count.example.item
}
