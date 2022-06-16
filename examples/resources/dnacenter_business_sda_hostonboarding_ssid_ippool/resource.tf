
resource "dnacenter_business_sda_hostonboarding_ssid_ippool" "example" {
  provider = dnacenter
  parameters {

    scalable_group_name = "string"
    site_name_hierarchy = "string"
    ssid_names          = ["string"]
    vlan_name           = "string"
  }
}

output "dnacenter_business_sda_hostonboarding_ssid_ippool_example" {
  value = dnacenter_business_sda_hostonboarding_ssid_ippool.example
}