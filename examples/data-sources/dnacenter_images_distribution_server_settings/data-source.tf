
data "dnacenter_images_distribution_server_settings" "example" {
  provider = dnacenter
}

output "dnacenter_images_distribution_server_settings_example" {
  value = data.dnacenter_images_distribution_server_settings.example.items
}
