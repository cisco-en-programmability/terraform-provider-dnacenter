
data "dnacenter_sda_fabric_sites_count" "example" {
  provider = dnacenter
}

output "dnacenter_sda_fabric_sites_count_example" {
  value = data.dnacenter_sda_fabric_sites_count.example.item
}
