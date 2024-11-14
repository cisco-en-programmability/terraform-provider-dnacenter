
resource "dnacenter_sites_dns_settings" "example" {
  provider = dnacenter

  parameters {

    dns {

      dns_servers = ["string"]
      domain_name = "string"
    }
    id = "string"
  }
}

output "dnacenter_sites_dns_settings_example" {
  value = dnacenter_sites_dns_settings.example
}