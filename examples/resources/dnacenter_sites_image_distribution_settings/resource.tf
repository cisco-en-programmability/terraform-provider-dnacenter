
resource "dnacenter_sites_image_distribution_settings" "example" {
  provider = dnacenter

  parameters {

    id = "string"
    image_distribution {

      servers = ["string"]
    }
  }
}

output "dnacenter_sites_image_distribution_settings_example" {
  value = dnacenter_sites_image_distribution_settings.example
}