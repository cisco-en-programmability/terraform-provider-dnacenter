
resource "dnacenter_sda_fabric_zones" "example" {
  provider = dnacenter

  parameters {

    authentication_profile_name = "string"
    id                          = "string"
    site_id                     = "string"
  }
}

output "dnacenter_sda_fabric_zones_example" {
  value = dnacenter_sda_fabric_zones.example
}