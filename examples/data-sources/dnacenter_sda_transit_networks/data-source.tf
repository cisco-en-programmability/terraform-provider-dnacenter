
data "dnacenter_sda_transit_networks" "example" {
  provider = dnacenter
  id       = "string"
  limit    = 1
  name     = "string"
  offset   = 1
  type     = "string"
}

output "dnacenter_sda_transit_networks_example" {
  value = data.dnacenter_sda_transit_networks.example.items
}
