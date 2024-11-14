
resource "dnacenter_sda_multicast_update" "example" {
  provider            = dnacenter
  site_name_hierarchy = "string"
}

output "dnacenter_sda_multicast_update_example" {
  value = dnacenter_sda_multicast_update.example
}