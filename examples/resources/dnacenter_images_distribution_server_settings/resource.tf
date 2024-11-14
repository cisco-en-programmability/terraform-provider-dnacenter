
resource "dnacenter_images_distribution_server_settings" "example" {
  provider = dnacenter

  parameters {

    password       = "******"
    port_number    = 1.0
    root_location  = "string"
    server_address = "string"
    username       = "string"
  }
}

output "dnacenter_images_distribution_server_settings_example" {
  value = dnacenter_images_distribution_server_settings.example
}