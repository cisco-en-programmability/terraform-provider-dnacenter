
resource "dnacenter_sites_ntp_settings" "example" {
  provider = dnacenter

  parameters {

    id = "string"
    ntp {

      servers = ["string"]
    }
  }
}

output "dnacenter_sites_ntp_settings_example" {
  value = dnacenter_sites_ntp_settings.example
}