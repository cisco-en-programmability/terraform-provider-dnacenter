
resource "dnacenter_sda_fabric_site" "example" {
  provider = dnacenter
  parameters {

    fabric_name         = "string"
    site_name_hierarchy = "string"
  }
}

output "dnacenter_sda_fabric_site_example" {
  value = dnacenter_sda_fabric_site.example
}