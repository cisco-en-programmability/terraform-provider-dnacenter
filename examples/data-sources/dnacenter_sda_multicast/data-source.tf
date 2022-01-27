
data "dnacenter_sda_multicast" "example" {
  provider            = dnacenter
  site_name_hierarchy = "string"
}

output "dnacenter_sda_multicast_example" {
  value = data.dnacenter_sda_multicast.example.item
}
