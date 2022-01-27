
data "dnacenter_business_sda_hostonboarding_ssid_ippool" "example" {
  provider            = dnacenter
  site_name_hierarchy = "string"
  vlan_name           = "string"
}

output "dnacenter_business_sda_hostonboarding_ssid_ippool_example" {
  value = data.dnacenter_business_sda_hostonboarding_ssid_ippool.example.item
}
