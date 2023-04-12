
data "dnacenter_business_sda_virtual_network_summary" "example" {
  provider            = dnacenter
  site_name_hierarchy = "string"
}

output "dnacenter_business_sda_virtual_network_summary_example" {
  value = data.dnacenter_business_sda_virtual_network_summary.example.item
}
