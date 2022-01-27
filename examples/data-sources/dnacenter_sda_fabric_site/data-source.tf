
data "dnacenter_sda_fabric_site" "example" {
  provider            = dnacenter
  site_name_hierarchy = "string"
}

output "dnacenter_sda_fabric_site_example" {
  value = data.dnacenter_sda_fabric_site.example.item
}
