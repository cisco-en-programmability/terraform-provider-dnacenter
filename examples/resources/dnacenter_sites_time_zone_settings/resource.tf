
resource "dnacenter_sites_time_zone_settings" "example" {
  provider = dnacenter

  parameters {

    id = "string"
    time_zone {

      identifier = "string"
    }
  }
}

output "dnacenter_sites_time_zone_settings_example" {
  value = dnacenter_sites_time_zone_settings.example
}