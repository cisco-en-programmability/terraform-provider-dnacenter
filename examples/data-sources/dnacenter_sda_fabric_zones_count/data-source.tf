
data "dnacenter_sda_fabric_zones_count" "example" {
  provider = dnacenter
}

output "dnacenter_sda_fabric_zones_count_example" {
  value = data.dnacenter_sda_fabric_zones_count.example.item
}
