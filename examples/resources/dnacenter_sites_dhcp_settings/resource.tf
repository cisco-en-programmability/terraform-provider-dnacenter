
resource "dnacenter_sites_dhcp_settings" "example" {
  provider = dnacenter

  parameters {

    dhcp {

      servers = ["string"]
    }
    id = "string"
  }
}

output "dnacenter_sites_dhcp_settings_example" {
  value = dnacenter_sites_dhcp_settings.example
}