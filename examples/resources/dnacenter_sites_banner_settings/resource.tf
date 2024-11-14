
resource "dnacenter_sites_banner_settings" "example" {
  provider = dnacenter

  parameters {

    banner {

      message = "string"
      type    = "string"
    }
    id = "string"
  }
}

output "dnacenter_sites_banner_settings_example" {
  value = dnacenter_sites_banner_settings.example
}