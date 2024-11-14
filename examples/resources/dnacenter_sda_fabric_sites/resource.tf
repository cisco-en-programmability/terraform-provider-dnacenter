
resource "dnacenter_sda_fabric_sites" "example" {
  provider = dnacenter

  parameters {

    authentication_profile_name = "string"
    id                          = "string"
    is_pub_sub_enabled          = "false"
    site_id                     = "string"
  }
}

output "dnacenter_sda_fabric_sites_example" {
  value = dnacenter_sda_fabric_sites.example
}